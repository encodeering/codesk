.PHONY: build setup

OUT := npiperelay.exe

build:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ${OUT} github.com/jstarks/npiperelay

setup:
	go get -d github.com/jstarks/npiperelay
