#!/bin/sh -e

go test -count=1 -race ./...
go vet ./...

go build -tags netgo ./cmd/usersvc/...
go build -tags netgo ./cmd/usercli/...
