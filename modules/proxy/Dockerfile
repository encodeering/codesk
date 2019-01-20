FROM       golang:1.11.4-alpine3.8
MAINTAINER encodeering@gmail.com

RUN        apk --no-cache add git gcc make libc-dev

WORKDIR    /usr/local/src
COPY       . .

RUN        make setup \
 &&        make test \
 &&        make build
