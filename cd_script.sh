#!/bin/sh -e

go test ./...

go build -tags netgo ./cmd/usersvc/...
docker build -t latanassov/usersvc:0.1.0 .

docker login
docker push latanassov/usersvc:0.1.0
# from here on openshift has to take over