#!/bin/bash
go build -i -o micro  ./vendor/github.com/micro/micro/main.go
./micro --server=grpc --client=grpc --transport=grpc --broker=grpc --registry=etcdv3  web


go build cmd/works/main.go

go build -o  run cmd/works/main.go

./run  --server_address=0.0.0.0:50001