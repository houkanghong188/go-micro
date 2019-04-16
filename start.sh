#!/bin/bash
go build -i -o micro  ./vendor/github.com/micro/micro/main.go
./micro --server=grpc --client=grpc --transport=grpc --broker=grpc --registry=etcdv3  web

go build -i -o works  cmd/works/main.go
./run  --server_address=0.0.0.0:50001

go build -i -o  template cmd/template/main.go
./template