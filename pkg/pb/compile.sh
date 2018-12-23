#!/bin/sh -e

protoc usersvc.proto --go_out=plugins=grpc:.