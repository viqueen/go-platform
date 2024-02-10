#!/usr/bin/env bash

function init() {
  name=${1}
  target=platform/${name}
  mkdir -p "${target}"
  go mod init -C "${target}" github.com/viqueen/go-platform/"${target}"
  go work use -r .
}

function tidy() {
  name=${1}
  target=platform/${name}
  go mod tidy -C "${target}"
}

eval "$@"