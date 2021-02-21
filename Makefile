# Go Parameters
GOCMD=go
GOBASE := $(shell pwd)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=grpclab

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

.PHONY: vendor
vendor:
	$(GOCMD) mod tidy && $(GOCMD) mod vendor

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: image
image:
	docker build -t user3301/grpclab .

.PHONE: spanner
spanner:
	docker run -p 9010:9010 -p 9020:9020 \
	--env SPANNER_PROJECT_ID=test \
	--env SPANNER_INSTANCE_ID=test-instance \
	--env SPANNER_DATABASE_ID=db \
	 --rm roryq/spanner-emulator:latest