### etcdv3 安装

### mac install protoc
    brew install protoc-gen-go
    go get -d -u github.com/golang/protobuf/protoc-gen-go

### install micro工具
    go install github.com/micro/micro
    如果是 报错包不存在，则可以使用下列命令可构建 
    go build -i -o micro  ./vendor/github.com/micro/micro/main.go 
    
    micro web 界面启动:
    ./micro --server=grpc --client=grpc --transport=grpc --broker=grpc --registry=etcdv3  web
    
    或者可以直接运行 sh start.sh 来运行 web 界面

### 使用 etcd3 服务发现
    启动时 如果 address 不是 127.0.0.1:2379 则需要加上 --registry_address
    service --registry=etcd --registry_address=127.0.0.1:2379
    
### 使用 protoc 生成代码 
    protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
### 问题列表

### api 调用
    1 启动 api 
    ./micro --server=grpc --client=grpc --transport=grpc --broker=grpc --registry=etcdv3  api
    2 调用 bank 服务 bank.show 
    HOST: http://localhost:8080/rpc
    HEADER: Content-Type: application/json
    BODY: {"service": "go.micro.srv.bank", "method": "Bank.Show", "request": {"id":1}}


### 生成gen -gorm
    1. 安装gen
     
    2. 命令运行
    gen --connstr "testgroup:testgroupM1@(rm-2zezj9n2lv83nl8x4o.mysql.rds.aliyuncs.com:3306)/makaplatv4?charset=utf8&parseTime=True&loc=Local" --table platv5_audit_conf --gorm