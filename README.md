# go-security-txt
A proof of concept program that pulls and parses security.txt files at mass.

## Description
A Go command-line tool that checks the existence of a `security.txt` file. Read the related blog post on: https://redmaple.tech/blogs/2022/survey-of-security-txt/ which has been published.

## Building
To build the tool from source you need a Go environment. I recommend using the latest Golang version available.

To support cross-platform builds, run:
```bash
GOOS=darwin GOARCH=arm64 go build -o sectxt-macos-arm64
GOOS=darwin GOARCH=amd64 go build -o sectxt-macos-intel
GOOS=windows GOARCH=amd64 go build -o sectxt-x64.exe
GOOS=linux GOARCH=amd64 go build -o sectxt-linux-x64
```

## Usage
The program can either take a domain name `-d`, or a text based input file `-i`, where each domain name is on a new line. You can also use `-json` option to display JSON formatted output, which is only supported with `-d` option.

```bash
naz@ThinkPad:~/dev/go/go-security-txt$ ./sectxt
  -d string
        specify a single domain name e.g. google.com
  -i string
        input filename of a list of domain names e.g. domains.txt
  -json
        output result to JSON format (only supports -d mode)
  -o string
        output filename of exported CSV file
```
Here is an example of using `-json` option and piping with `jq` to make it pretty.
```bash
naz@Naz-ThinkPad:~/dev/go/go-security-txt$ ./sectxt -json -d redmaple.tech | jq
{
  "Name": "redmaple.tech",
  "IsFileFound": true,
  "IsFileValid": true,
  "Result": {
    "Acknowledgments": "",
    "Canonical": "https://redmaple.tech/.well-known/security.txt",
    "Contact": [
      "mailto:security@redmaple.tech",
      "https://redmaple.tech/contact/"
    ],
    "Encryption": "Talk to us about https://redmaple.tech/products/#trebuchet",
    "Expires": "2023-05-30T23:00:00.000Z",
    "Hiring": "",
    "Policy": "https://redmaple.tech/disclosure/",
    "PreferredLanguages": [
      "en"
    ],
    "CSAF": null
  },
  "Complete": true,
  "StatusCode": "200",
  "Error": ""
}
```

Here is an example of using `-i` input file, and `-o` which will create a CSV formatted result file.
```bash
naz@Naz-ThinkPad:~/dev/go/go-security-txt$ ./sectxt -i datasets/uk-banks.txt -o out/test.csv
2022/08/18 17:12:02 Read 20 domains from input file
2022/08/18 17:12:02 Creating 500 workers
2022/08/18 17:12:02 Running...
2022/08/18 17:12:03 Writer done, wrote 20 rows
2022/08/18 17:12:03 All done.
```

The CSV output will look like this:
```csv
domain,is_file_found,is_file_valid,http_status,acknowledgments,canonical,contact,encryption,expires,hiring,policy,preferred_languages,errors
monzo.com,true,true,200,,,"https://hackerone.com/monzo, mailto:security@monzo.com",https://monzo.com/.well-known/monzo-publickey.asc,,https://monzo.com/careers/,,,
www.bankofscotland.co.uk,false,false,200,,,,,,,,,Content-type (text/html) is not valid
www.revolut.com,true,true,200,,,"https://www.revolut.com/responsible-disclosure-program, security@revolut.com",,,,,,
www.halifax.co.uk,false,false,200,,,,,,,,,Content-type (text/html) is not valid
www.lloydsbank.com,false,false,200,,,,,,,,,Content-type (text/html) is not valid
www.nationwide.co.uk,true,true,200,,,mailto:securitydisclosure@nationwide.co.uk,,2023-01-25T17:00:00.000Z,,,en,
www.natwest.com,false,false,404,,,,,,,,,https://www.natwest.com/security.txt file may not exist - HTTP error code: 404
```

## Results
All the survey results can be found within the `out` folder. These include:
- Tranco 1m
- Moz 500
- FTSE 100
- S&P 500
- UK financial industry
- UK Banks
