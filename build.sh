#!/bin/bash

ARCH="amd64"

if [ $# -eq 0 ]; then
  echo "linux"
  OS="linux"
else
  # darwin / linux / windows
  echo "$1"
  OS="$1"
fi

GOOS=$OS GOARCH=$ARCH go build -o at-once cmd/at-once.go
GOOS=$OS GOARCH=$ARCH go build -o per-sec cmd/per-sec.go
