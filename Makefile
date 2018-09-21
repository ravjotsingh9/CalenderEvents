include ./testData/Makefile

.DEFAULT_GOAL := binary

deps: gendeps
	go get github.com/tatsushid/go-prettytable
	
binary: 
	go build -o ecf_binary main.go 
	-@chmod +x ecf_binary

	