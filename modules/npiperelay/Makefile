.PHONY: build setup

build:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o /usr/bin/wsl-npiperelay.exe github.com/jstarks/npiperelay

setup:
	go get -d github.com/jstarks/npiperelay
