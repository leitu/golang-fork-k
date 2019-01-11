# Go parameters
BINARY_NAME=k
BINARY_LUNIX=$(BINARY_NAME)_linux

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

GO111MODULE=on


all: build

build:
	$(GOBUILD) -o $(BINARY_NAME)

# TODO: Test
#test: 
#    $(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LUNIX)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LUNIX)