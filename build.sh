#!/bin/bash

cd bin

# 构建 linux template 服务
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ../cmd/template/main.go
mv main template
echo "SUCCESS TEMPLATE"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ../cmd/font/main.go
mv main font
echo "SUCCESS FONT"

