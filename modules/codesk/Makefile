.PHONY: all build clean generate test setup
.DEFAULT_GOAL := all

OUT := codesk.exe

all: clean test build

clean:
	rm ${OUT} 2> /dev/null || true
	packr2 clean

generate:
	go generate
	GO111MODULE=on packr2

test: generate
	$(MAKE) -C glue test
	go test ./internal/...

build: generate
	GOOS=windows GOARCH=amd64 GO111MODULE=on packr2 build -ldflags "-s -w" -o ${OUT}

setup:
	$(MAKE) -C glue setup
	go get -u github.com/gobuffalo/packr/v2/packr2
	go get
