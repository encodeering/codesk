FROM       golang:1.11.4-alpine3.8 as builder
MAINTAINER encodeering@gmail.com

RUN        apk --no-cache add git gcc make libc-dev

WORKDIR    /usr/local/src
COPY       . .

ARG        TARGET

RUN        make test \
 &&        make build TARGET=${TARGET}



FROM       scratch

ARG        TARGET

COPY       --from=builder /usr/local/src/*.exe /
