package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FileExists checks the folder exists and is a folder, return boolean result
func FileExists(filePath string) bool {
	stat, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	if stat.IsDir() {
		return false
	}

	return true
}

// ReadFile lines into a slice of domain names
func ReadFile(f string) ([]string, error) {

	fileHandle, err := os.Open(f)

	if err != nil {
		return nil, fmt.Errorf("Cannot open file: %s", err.Error())
	}

	// Close handle
	defer fileHandle.Close()

	domains := []string{}
	data := bufio.NewScanner(fileHandle)

	for data.Scan() {

		line := data.Text()
		domain := strings.TrimSpace(line)
		// For each line extract the top-level domain
		if !IsValidDomain(domain) {
			fmt.Printf("[!] %s is a invalid domain\n", domain)
			// ignore malformed domain names
			continue
		}
		// add valid domains to slice
		domains = append(domains, domain)
	}

	return domains, nil
}
