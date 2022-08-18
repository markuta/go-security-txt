package request

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/markuta/go-security-txt/parser"
)

func handleRequest(domain string) (r []byte, statusCode int, err error) {

	url := []string{
		"https://" + domain + securityTXT,
		"https://" + domain + securityTXTAlt,
	}

	req, _ := http.NewRequest("GET", url[0], nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "close")
	//if err != nil {
	//	return nil, "", fmt.Errorf("HTTP request failed: %s", err.Error())
	//}

	//ctx, cancel := context.WithTimeout(context.Background(), HTTPtimeoutSecs*time.Second)
	//defer cancel()

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
				//KeepAlive: 30 * time.Second,
			}).Dial,
			// Avoid: "x509: certificate signed by unknown authority"
			TLSClientConfig: &tls.Config{
				//InsecureSkipVerify: true,
			},
			ForceAttemptHTTP2: true,
			IdleConnTimeout:   2 * time.Second,
		},
	}

	//req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("HTTP response error: %s", err)
	}
	// Close handle
	defer res.Body.Close()

	// Try with alternative PATH
	if res.StatusCode == 404 || res.StatusCode == 403 {
		req, _ = http.NewRequest("GET", url[1], nil)
		res, err = client.Do(req)
		if err != nil {
			return nil, 0, fmt.Errorf("HTTP response error: %s", err)
		}
		// Close handle
		defer res.Body.Close()
	}

	// Accept HTTP 2xx responses with content type text/plain
	if !isHTTPResponseValid(res.StatusCode) {
		return nil, res.StatusCode, fmt.Errorf("%s %s%d", res.Request.URL, HTTPError, res.StatusCode)
	} else if !(strings.HasPrefix(res.Header.Get("Content-type"), "text/plain")) {
		return nil, res.StatusCode, fmt.Errorf("Content-type (%s) is not valid", res.Header.Get("Content-type"))
	}

	// Fetch response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, fmt.Errorf("Cannot read response body: %s", err)
	}

	return data, res.StatusCode, nil
}

// Process performs the calls to request functions
func Process(d string) (*parser.Domain, error) {

	domain := parser.Domain{Name: d}

	body, status, err := handleRequest(d)
	domain.StatusCode = strconv.Itoa(status)

	if err != nil {
		return &domain, err
	}

	domain.IsFileFound = true

	// Check body size
	if len(body) < 5 {
		return &domain, fmt.Errorf("File is empty")
	}

	// Parse and extract data from security.txt
	// Store within the Domain.Result struct
	secTxtPtr := parser.ParseSecTXT(body)
	domain.Result = *secTxtPtr

	if !checkAllFieldsEmpty(domain) {
		domain.IsFileValid = true
	}

	return &domain, nil
}

// A really ugly way to check if struct is empty
func checkAllFieldsEmpty(domain parser.Domain) bool {
	return domain.Result.Acknowledgments == "" && domain.Result.Contact == nil && domain.Result.Encryption == "" &&
		domain.Result.Expires == "" && domain.Result.Hiring == "" && domain.Result.Policy == "" &&
		domain.Result.PreferredLanguages == nil
}

// checkHTTPResponse()
func isHTTPResponseValid(statusCode int) bool {
	return statusCode >= statusOK && statusCode <= statusIMUsed
}
