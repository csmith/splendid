#!/bin/sh

set -eux

go tool golang.org/x/tools/cmd/goimports -w .
go tool honnef.co/go/tools/cmd/staticcheck ./...
go test ./...
