.PHONY: build test

OUT := codesk-$(TARGET).exe

test:
	go test ./internal/...

build:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.target=$(TARGET)" -o ${OUT}
