.DEFAULT_GOAL := run

.PHONY:fmt vet build
fmt:
				go fmt .

vet: fmt 
				go vet .

build: vet
				go build .

clean: 
				go clean .

run:
				go run . file.txt