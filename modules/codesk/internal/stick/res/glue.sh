#!/usr/bin/env bash

shopt -s expand_aliases

alias docker=docker.exe

set -x -euo pipefail

{
    target="${1}"
    output="$(basename ${target}).exe"
    tag="codesk-proxy:${output}"

    docker build -t "${tag}" --build-arg TARGET="${target}" 1>&2 -
} > /dev/null

cid=$(docker create "${tag}" true)
docker cp "${cid}:${output}" - | tar xOf - "${output}"

{
    docker rm -v "${cid}" || true
} > /dev/null
