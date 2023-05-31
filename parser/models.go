package parser

// SecTXTFile struct for security.txt file
type SecTXTFile struct {
	//Comments        string // e.g. # This is a comment
	Acknowledgments    string   // e.g. https://example.com/halloffame
	Canonical          string   // e.g. https://example.com/.well-known/security.txt
	Contact            []string // e.g. https://example.com/security or mailto:security@example.com
	Encryption         string   // e.g. https://example.com/publickey.txt
	Expires            string   // e.g. Expires: 2021-12-31T18:37:07z
	Hiring             string   // e.g. https://example.com/careers
	Policy             string   // e.g. https://example.com/disclosure.html
	PreferredLanguages []string // e.g. en, es, fr
	CSAF               []string // e.g. https://example.com/.well-known/csaf/provider-metadata.json
}

// Domain struct shows whether a domain
// name has a security.txt file or not. If
// it does, include the Result struct too.
type Domain struct {
	Name        string // e.g. example.com
	IsFileFound bool
	IsFieldFound bool
	Result      SecTXTFile
	Complete    bool
	StatusCode  string
	Error       string
}
