package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"text/tabwriter"
)

// PrettyPrint is a method for calling on declared structs
func (d Domain) PrettyPrint() {
	// Print status
	w := new(tabwriter.Writer)
	w.Init(log.Writer(), 8, 8, 2, '\t', 0)

	fmt.Fprintf(w, "Domain:\t%s\n", d.Name)
	fmt.Fprintf(w, "Acknowledgments:\t%s\n", d.Result.Acknowledgments)
	fmt.Fprintf(w, "Canonical:\t%s\n", d.Result.Canonical)
	fmt.Fprintf(w, "Contact:\t%s\n", SliceAsCSV(d.Result.Contact))
	fmt.Fprintf(w, "Encryption:\t%s\n", d.Result.Encryption)
	fmt.Fprintf(w, "Expires\t%s\n", d.Result.Expires)
	fmt.Fprintf(w, "Hiring\t%s\n", d.Result.Hiring)
	fmt.Fprintf(w, "Policy\t%s\n", d.Result.Policy)
	fmt.Fprintf(w, "PreferredLanguages\t%s\n", SliceAsCSV(d.Result.PreferredLanguages))

	w.Flush()
}

// SliceAsCSV converts slice to string separated by ,
func SliceAsCSV(slice []string) string {
	return strings.Join(slice, ", ")
}

// JSONExport transforms SecTXTFile into JSON
func (d Domain) JSONExport() ([]byte, error) {
	dataJSON, err := json.Marshal(d)
	return dataJSON, err
}

// CSVExport exports SecTXTFile into a CSV file
//func (s SecTXTFile) CSVExport() ([]byte, error) {
//	return dataCSV, err
//}
