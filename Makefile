    # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
MAIN= cmd/main.go
BINARY_NAME=fuego-cache


all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)

test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)
	./$(BINARY_NAME)

fmt:
	$(GOFMT) -w .

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

docker-build-cli:
	@docker build -t fuego-cache-cli .

docker-build-http:
	@docker build -t fuego-cache-http .

docker-run-cli:
	@docker run -it -p 9919:9919 fuego-cache-cli

docker-run-http:
	@docker run -d -p 9999:9999 fuego-cache-http
