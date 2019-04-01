#!/bin/bash
go build -i -o micro  ./vendor/github.com/micro/micro/main.go
./micro --server=grpc --client=grpc --transport=grpc --broker=grpc --registry=etcdv3  web