CMD = snyk-to-md
BIN = bin/$(CMD)
BIN_DARWIN = $(BIN)-darwin
BIN_LINUX = $(BIN)-linux
SOURCES = $(shell find . -type f -iname "*.go")
CMD_SRC = ./main.go

.PHONY: all build vet fmt clean

$(BIN_DARWIN): $(SOURCES)
	GOARCH=arm64 GOOS=darwin go build -o $(BIN_DARWIN) $(CMD_SRC)

$(BIN_LINUX): $(SOURCES)
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o $(BIN_LINUX) $(CMD_SRC)

build: $(BIN_DARWIN) $(BIN_LINUX)

vet:
	go vet ./...

fmt:
	go fmt ./...

run: fmt vet build
	${BIN_DARWIN}

test:
	go test ./...

clean:
	rm -rf bin/

test-input: fmt vet build
	${BIN_DARWIN} -i test-data/dummy-report.json
