# Go Parameters
GOCMD=go
GOBASE := $(shell pwd)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
COMMITSHA=$(shell git rev-parse --short HEAD)
BINARY_NAME=goate_$(COMMITSHA)

all: test build

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) $(GOBASE)/cmd/.

.PHONY: test
test:
	$(GOTEST) -v -cover ./...

.PHONY: clean
clean:
	$(GOCLEAN) ./...
	rm -f $(BINARY_NAME)

.PHONY: rebuild
rebuild: clean build

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: image
image:
	docker build -t user3301/goate .

.PHONY: gen-proto
gen-proto:
	docker run -rm -u
