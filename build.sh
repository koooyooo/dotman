#!/bin/bash

ARCH="amd64"

if [ $# -eq 0 ]; then
  OS="linux"
else
  # darwin / linux / windows
  OS="$1"
fi

GOOS=$OS GOARCH=$ARCH go build -o at-once cmd/at-once.go
GOOS=$OS GOARCH=$ARCH go build -o per-sec cmd/per-sec.go
