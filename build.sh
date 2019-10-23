#!/bin/bash

export GOARCH="amd64"

if [ $# -eq 0 ]; then
  export GOOS="linux"
else
  # darwin / linux / windows
  export GOOS="$1"
fi

go build -o at-once cmd/at-once.go
go build -o per-sec cmd/per-sec.go
