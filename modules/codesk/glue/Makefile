.PHONY: all build clean test
.DEFAULT_GOAL := all

OUT := $(TARGET).exe

all: clean test build

clean:
	rm ${OUT} 2> /dev/null || true

test:
	go test ./internal/...

build:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.target=$(TARGET)" -o ${OUT}
