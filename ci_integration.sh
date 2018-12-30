#!/bin/sh

CUR_DIR=$(pwd)


CONTAINER_ID=$(docker run -d -p 8080:8080 latanassov/usersvc:0.1.0)

mkdir -p report

go test -race -coverprofile=./report/coverage_it.out  -tags=integration ./... 
go tool cover -html=./report/coverage_it.out -o ./report/coverage_it.html

docker stop $CONTAINER_ID >/dev/null