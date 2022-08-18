package test

import (
	"testing"

	"github.com/markuta/go-security-txt/request"
	"github.com/markuta/go-security-txt/utils"
)

const (
	rmtDomain       = "redmaple.tech"
	malformedDomain = " redmaple.tech " // contains spaces
	rmtCanonical    = "https://redmaple.tech/.well-known/security.txt"
)

func TestIsValidDomain(t *testing.T) {
	if !utils.IsValidDomain(malformedDomain) {
		t.Logf("%q is not valid", malformedDomain)
	} else {
		t.Errorf("This domain is not valid  %q\n", malformedDomain)
	}

}

func TestValid(t *testing.T) {
	_, err := request.Process(rmtDomain)

	if err != nil {
		t.Errorf("Error getting result for %q: %s\n", rmtDomain, err.Error())
	}
}

func TestStruct(t *testing.T) {
	domainData, err := request.Process(rmtDomain)

	if err != nil {
		t.Errorf("Error getting result for %q: %s\n", rmtDomain, err.Error())
	}

	if domainData.Result.Canonical != rmtCanonical {
		t.Errorf("secTXT.Canonical is %q, not %q\n", domainData.Result.Canonical, rmtCanonical)
	}
}
