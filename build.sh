#!/bin/bash
go build -i -o micro  ./vendor/github.com/micro/micro/main.go
go build -i -o works  cmd/works/main.go
go build -i -o  template cmd/template/main.go
