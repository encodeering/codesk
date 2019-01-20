.PHONY: build test

test:
	go test ./internal/...

build:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.target=$(TARGET)" -o /usr/bin/codesk-$(TARGET).exe
