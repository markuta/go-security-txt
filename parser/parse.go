package parser

import (
	"bufio"
	"bytes"
	"strings"
)

/*
	https://www.rfc-editor.org/rfc/rfc9116
 	https://datatracker.ietf.org/doc/html/rfc9116

     2.5.  Field Definitions
       2.5.1.  Acknowledgments
       2.5.2.  Canonical
       2.5.3.  Contact
       2.5.4.  Encryption
       2.5.5.  Expires
       2.5.6.  Hiring
       2.5.7.  Policy
       2.5.8.  Preferred-Languages

*/
/*
// ParseSecTXT uses reflection to get field names
func ParseSecTXT(r []byte) *SecTXTFile {

	// Static prefix
	prefix := map[string]string{
		"Acknowledgments":    "Acknowledgements: ",
		"Contact":            "Contact: ",
		"Canonical":          "Canonical: ",
		"Encryption":         "Encryption: ",
		"Expires":            "Expires: ",
		"Hiring":             "Hiring: ",
		"Policy":             "Policy: ",
		"PreferredLanguages": "Preferred-Languages: ",
	}

	secTXT := SecTXTFile{}
	contact := []string{} // Contact may be on multiple lines

	lines := getLines(r)

	// Parse each line
	for _, rawLine := range lines {
		// Iterate through prefix slice
		line := strings.TrimSpace(rawLine)

		for fieldName, prefixStr := range prefix {
			// If beginning of line matches prefix
			if strings.Contains(line, prefixStr) {
				// Use reflection to match struct and assign string value to field
				valueStr := strings.Split(line, prefixStr)
				// Assign string to the matching field
				field := reflect.ValueOf(&secTXT).Elem().FieldByName(fieldName)
				// Store two fields as []string slices
				if fieldName == "Contact" {
					contact = append(contact, valueStr[1])
					field.Set(reflect.ValueOf(contact))
				} else if fieldName == "PreferredLanguages" {
					s := strings.Split(valueStr[1], ",")
					field.Set(reflect.ValueOf(s))
				} else {
					// Store everything else as string
					field.SetString(valueStr[1])
				}

			}

		}
	}

	return &secTXT
} */

// ParseSecTXT using switch statement instead
func ParseSecTXT(r []byte) *SecTXTFile {

	// Static prefix
	prefix := map[string]string{
		"Acknowledgments":    "Acknowledgements: ",
		"Contact":            "Contact: ",
		"Canonical":          "Canonical: ",
		"Encryption":         "Encryption: ",
		"Expires":            "Expires: ",
		"Hiring":             "Hiring: ",
		"Policy":             "Policy: ",
		"PreferredLanguages": "Preferred-Languages: ",
	}

	secTXT := SecTXTFile{}
	//data := NormalizeNewlines(r)
	//lines := strings.Split(string(data), "\n")

	rawLines := getLines(r)

	// Parse each line
	for _, rawLine := range rawLines {
		// Iterate through prefix slice
		line := strings.TrimSpace(rawLine)

		switch {
		case strings.HasPrefix(line, prefix["Acknowledgments"]):
			valueStr := strings.Split(line, prefix["Acknowledgments"])
			secTXT.Acknowledgments = valueStr[1]
		case strings.HasPrefix(line, prefix["Contact"]):
			valueStr := strings.Split(line, prefix["Contact"])
			// Contact may be on multiple lines
			secTXT.Contact = append(secTXT.Contact, valueStr[1])
		case strings.HasPrefix(line, prefix["Canonical"]):
			valueStr := strings.Split(line, prefix["Canonical"])
			secTXT.Canonical = valueStr[1]
		case strings.HasPrefix(line, prefix["Encryption"]):
			valueStr := strings.Split(line, prefix["Encryption"])
			secTXT.Encryption = valueStr[1]
		case strings.HasPrefix(line, prefix["Expires"]):
			valueStr := strings.Split(line, prefix["Expires"])
			secTXT.Expires = valueStr[1]
		case strings.HasPrefix(line, prefix["Hiring"]):
			valueStr := strings.Split(line, prefix["Hiring"])
			secTXT.Hiring = valueStr[1]
		case strings.HasPrefix(line, prefix["Policy"]):
			valueStr := strings.Split(line, prefix["Policy"])
			secTXT.Policy = valueStr[1]
		case strings.HasPrefix(line, prefix["PreferredLanguages"]):
			valueStr := strings.Split(line, prefix["PreferredLanguages"])
			langStr := strings.Split(valueStr[1], ",")
			secTXT.PreferredLanguages = langStr
		}
	}

	return &secTXT
}

func getLines(d []byte) []string {

	var lines []string

	bytesData := bytes.NewReader(normalizeNewlines(d))
	scanner := bufio.NewScanner(bytesData)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// normalizeNewlines normalizes \r\n (windows) and \r (mac) into \n (unix)
// src: https://www.programming-books.io/essential/go/normalize-newlines-1d3abcf6f17c4186bb9617fa14074e48
func normalizeNewlines(d []byte) []byte {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return d
}

/* func getURLPath(line string) string {
	u, err := url.Parse(line)
	fmt.Println(u)
	if err != nil {
		log.Fatalf("Error extracting data%s", err.Error())
	}

	return u.Path
} */
