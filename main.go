package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/markuta/go-security-txt/parser"
	"github.com/markuta/go-security-txt/request"
	"github.com/markuta/go-security-txt/utils"
)

const (
	defaultWorkers = 50
)

var domain, inputFile, outputFilename string
var numWorkers int
var outputJSON bool

func init() {
	flag.StringVar(&domain, "d", "", "specify a single domain name e.g. google.com")
	flag.StringVar(&inputFile, "i", "", "input filename of a list of domain names e.g. domains.txt")
	flag.StringVar(&outputFilename, "o", "", "output filename of exported CSV file")
	flag.IntVar(&numWorkers, "w", defaultWorkers, "number of Goroutine workers")
	flag.BoolVar(&outputJSON, "json", false, "output result to JSON format (only supports -d mode)")
	//flag.BoolVar(&outputCSV, "csv", false, "(optional) output results to CSV format")
	//flag.BoolVar(&verboseErr, "v", false, "verbose logging")
	flag.Parse()
}

func main() {
	// single domain
	if domain != "" {
		runWithInputDomain()
	} else if inputFile != "" {
		// input file
		runWithInputFile()
	} else {
		flag.Usage()
		log.Fatalf("Missing domain input or file entered\n")
	}
}

func runWithInputFile() {
	// Simple file checks
	if outputFilename == "" {
		log.Fatalln("You must specify a output file (-o) when using the (-i) option")
	}

	if !utils.FileExists(inputFile) {
		log.Fatalf("Input Filename %s does not exist or you do not have permission to view it", inputFile)
	}

	domains, err := utils.ReadFile(inputFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Read %d domains from input file\n", len(domains))

	// Create CSV output file
	csvWriter, csvFile := utils.GetCSVWriter(outputFilename)

	// Write headers
	csvWriter.Write([]string{"timestamp","domain", "is_file_found", "is_field_found", "http_status", "acknowledgments", "canonical", "contact", "encryption", "expires", "hiring", "policy", "preferred_languages", "csaf", "errors"})
	numDomains := len(domains)

	// Catch CTRL-C so we can exit cleanly
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		utils.CloseWriter(csvWriter, csvFile)
		log.Println("Caught CTRL-C, exiting")
		os.Exit(1)
	}()

	// process all the domains
	processMultipleDomains(numDomains, domains, csvWriter)

	// flush and close
	utils.CloseWriter(csvWriter, csvFile)
}

func runWithInputDomain() {
	// Use basic domain validator
	if !utils.IsValidDomain(domain) {
		log.Fatalf("%s doesn't look like a domain", domain)
	}

	result, err := processSingleDomain(domain)

	if err != nil {
		// check for HTTP error aka file not there

		if outputJSON {
			dataJSON, _ := result.JSONExport()
			fmt.Println(string(dataJSON))
		} else {
			fmt.Printf("%s doesn't have a security.txt file.\n", domain)
		}
		os.Exit(1)

	} else if outputJSON {
		dataJSON, _ := result.JSONExport()
		fmt.Println(string(dataJSON))
	} else {
		result.PrettyPrint()
	}
}

func processSingleDomain(domain string) (*parser.Domain, error) {
	domainResult, err := request.Process(domain)

	if err != nil {
		domainResult.Error = err.Error()
		return domainResult, err
	}

	domainResult.Complete = true

	return domainResult, nil
}

func processMultipleDomains(numDomains int, domains []string, csvWriter *csv.Writer) {
	// Setup channels
	resultsChan := make(chan *parser.Domain, numDomains)
	done := make(chan bool)

	// push jobs
	jobsChan := make(chan string, numDomains)

	for _, domain := range domains {
		jobsChan <- domain
	}

	// all the jobs sent, close the channel
	close(jobsChan)

	// Setup Results Writer
	go utils.CSVWriterRoutine(resultsChan, done, numDomains, csvWriter)

	// create workers
	log.Printf("Creating %d workers\n", numWorkers)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobsChan, resultsChan)
	}

	log.Println("Running...")

	// Notify main goroutine process is finished
	<-done

	// all workers done so close the results
	close(resultsChan)

	log.Printf("All done.")
}

func worker(id int, jobsChan <-chan string, resultsChan chan *parser.Domain) {
	for domain := range jobsChan {
		domainResult, err := request.Process(domain)

		if err != nil {
			domainResult.Error = err.Error() // Error as string
		}

		domainResult.Complete = true

		// Send data to channel
		resultsChan <- domainResult
	}
}
