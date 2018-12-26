#!/bin/sh

CUR_DIR=$(pwd)


CONTAINER_ID=$(docker run -d -p 8080:8080 latanassov/usersvc:0.1.0)

mkdir -p report

go test -count=1 -tags=integration ./... -coverprofile=./report/coverage_it.out -race
go tool cover -html=./report/coverage_it.out -o ./report/coverage_it.html

docker stop $CONTAINER_ID >/dev/null