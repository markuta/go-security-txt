package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const regEx = `^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`

// IsValidDomainCrap is very basic domain validator that may miss valid domains
func IsValidDomainCrap(domain string) bool {
	// source: https://www.socketloop.com/tutorials/golang-use-regular-expression-to-validate-domain-name
	regExp := regexp.MustCompile(regEx)
	return regExp.MatchString(domain)
}

// IsValidDomain is a better version of above
// source: https://ao.ms/how-to-create-a-domain-name-validator-in-python/
func IsValidDomain(domain string) bool {
	// Catch invalid domains first
	if len(domain) > 253 || len(domain) == 0 {
		//fmt.Println(1)
		return false
	}

	splitDomain := strings.Split(domain, ".")
	if len(splitDomain) > 127 || len(splitDomain) < 2 {
		//fmt.Println(2)
		return false
	}

	for _, x := range splitDomain {
		if len(x) > 63 || len(x) == 0 {
			//fmt.Println(3)
			return false
		}

		if !isAlphanumeric(rune(x[0])) || !isAlphanumeric(rune(x[len(x)-1])) {
			fmt.Printf(string(x[len(x)-1]))
			//fmt.Println(4)
			return false
		}

		// Check for valid unicode domain name characters
		// Link: https://en.wikipedia.org/wiki/List_of_Unicode_characters
		for _, l := range x {

			if !(int(l) < 128) || !isAlphanumeric(l) && l != '-' {
				fmt.Println(5)
				return false
			}
		}
	}

	tld := splitDomain[len(splitDomain)-1]
	// TLD shouldn't have any numbers
	if _, err := strconv.Atoi(tld); err == nil {
		// fmt.Println(6)
		return false
	}

	return true
}

func isAlphanumeric(l rune) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(string(l))
}
