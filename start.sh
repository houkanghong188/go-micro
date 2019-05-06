#!/bin/bash

lsof -i:8082 | awk '{print $2}' | xargs kill -9
lsof -i:50001 | awk '{print $2}' | xargs kill -9
lsof -i:50002 | awk '{print $2}' | xargs kill -9

cp ./config-example/config-eaxmple.go  ./config/config-eaxmple.go
go build -i -o micro  ./vendor/github.com/micro/micro/main.go
 nohup ./micro --server=grpc --client=grpc --transport=grpc --broker=grpc --registry=etcdv3  web &

go build -i -o works  cmd/works/main.go
nohup ./works  --server_address=0.0.0.0:50001 &

go build -i -o  template cmd/template/main.go &
nohup ./template --server_address=0.0.0.0:50002 &