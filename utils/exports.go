package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/markuta/go-security-txt/parser"
)

// GetCSVWriter create a handler for csv files
func GetCSVWriter(filename string) (*csv.Writer, *os.File) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	w := csv.NewWriter(f)
	w.UseCRLF = true
	return w, f
}

func CSVWriterRoutine(domainChannel chan *parser.Domain, done chan bool, numRecords int, csvWriter *csv.Writer) {
	rowsWritten := 0

	// Write data from channel to CSV
	for data := range domainChannel {
		err := csvWriter.Write([]string{
			data.Name,
			strconv.FormatBool(data.IsFileFound),
			strconv.FormatBool(data.IsFileValid),
			data.StatusCode,
			data.Result.Acknowledgments,
			data.Result.Canonical,
			parser.SliceAsCSV(data.Result.Contact),
			data.Result.Encryption,
			data.Result.Expires,
			data.Result.Hiring,
			data.Result.Policy,
			parser.SliceAsCSV(data.Result.PreferredLanguages),
			data.Error, // show general errors,
		})
		if err != nil {
			log.Fatalln("File writing failed")
			log.Fatalln(err)
		}

		rowsWritten++
		numRecords--

		// Check if all records are processed, if yes then notify channel
		if numRecords == 0 {
			done <- true
			log.Printf("Writer done, wrote %d rows\n", rowsWritten)
		}
	}
}

func CloseWriter(csvWriter *csv.Writer, csvFile *os.File) {
	// Flush buffer to file
	csvWriter.Flush()

	// Close the file
	err := csvFile.Close()

	if err != nil {
		log.Fatalf("Error closing file: %s", err.Error())
	}
}
