.PHONY: all build clean test setup
.DEFAULT_GOAL := all

OUT := $(shell basename $(TARGET)).exe

all: clean test build

clean:
	rm ${OUT} 2> /dev/null || true

test:
	go test ./internal/...

build:
ifdef TARGET
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.target=$(TARGET)" -o ${OUT}
endif

setup:
	go get
