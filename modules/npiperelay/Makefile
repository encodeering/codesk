.PHONY: all build clean setup
.DEFAULT_GOAL := all

OUT := npiperelay.exe

all: clean build

clean:
	rm ${OUT} 2> /dev/null || true

build:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ${OUT} github.com/jstarks/npiperelay

setup:
	go get -d github.com/jstarks/npiperelay
