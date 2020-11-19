all: build test

build:
	go build cmd/arrai-browse.go

install:
	go install cmd/arrai-browse.go

test:
	go test ./...

start:
	go run cmd/arrai-browse.go

tidy:
	go mod tidy
	gofmt -s -w .
	goimports -w .
