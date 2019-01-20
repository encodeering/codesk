.PHONY: build clean generate test setup

clean:
	packr2 clean

generate:
	go generate
	GO111MODULE=on packr2

test: generate
	$(MAKE) -C glue
	go test ./internal/...

build: generate
	GOOS=windows GOARCH=amd64 GO111MODULE=on packr2 build -ldflags "-s -w" -o /usr/bin/codesk.exe

setup:
	go get -u github.com/gobuffalo/packr/v2/packr2
