include ./testData/Makefile

.DEFAULT_GOAL := binary

deps: gendeps
	go get github.com/tatsushid/go-prettytable
	go get golang.org/x/tools/cmd/cover
	
binary: 
	go build -o ecf_binary main.go 
	-@chmod +x ecf_binary

test:
	go test -v

coverage:
	go test -coverprofile fmt
	