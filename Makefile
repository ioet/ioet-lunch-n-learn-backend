deps:
	go mod tidy
	go mod download

update:
	go get -u ./...

start:
	$(shell go env GOPATH)/bin/air

help:
	@echo "Available commands:"
	@echo "  make deps   - Download and install all Go dependencies"
	@echo "  make update - Update Go dependencies"
	@echo "  make start  - Start the project using air"
	@echo "  make help   - Show this help message"
