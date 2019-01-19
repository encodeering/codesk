FROM       golang:1.11.4-alpine3.8
MAINTAINER encodeering@gmail.com

RUN        apk --no-cache add git gcc libc-dev

WORKDIR    /usr/local/src
COPY       . .

ARG        TARGET

RUN                                  go test ./internal/... \
 &&        GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.target=${TARGET}" -o /usr/bin/wsl-${TARGET}.exe
