#!/usr/bin/env bash

shopt -s expand_aliases

alias docker=docker.exe

set -x -euo pipefail

{
    cwd=`dirname "${BASH_SOURCE}"`

    target="${1}"
    output="wsl-${target}.exe"
    tag="wsl-proxy:${target}"

    docker build -t "${tag}" --build-arg TARGET="${target}" "${cwd}"
} > /dev/null

cid=$(docker create "${tag}")
docker cp "${cid}:/usr/bin/${output}" - | tar xOf - "${output}"

{
    docker rm -v "${cid}" || true
} > /dev/null
