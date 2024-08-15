deps:
	go mod tidy
	go mod download

update:
	go get -u ./...

start:
	$(shell go env GOPATH)/bin/air

build:
	go build -v ./...

check-lint:
	$(shell go env GOPATH)/bin/golangci-lint run

lint:
	$(shell go env GOPATH)/bin/golangci-lint run --fix

test:
	echo "Not testing yet :("

help:
	@echo "Available commands:"
	@echo "  make deps       - Download and install all Go dependencies"
	@echo "  make update     - Update Go dependencies"
	@echo "  make start      - Start the project using air"
	@echo "  make build      - Build the project"
	@echo "  make help       - Show this help message"
	@echo "  make check-lint - Run linter"
	@echo "  make lint       - Run linter and fix"
	@echo "  make test       - Run tests"
