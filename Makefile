default:
	go build -o sectxt *.go

test:
	go test -v ./tests/...