### install micro工具
    go install github.com/micro/micro
### micro web 界面启动
    micro --server=grpc --client=grpc --transport=grpc --broker=grpc web    

### 使用 etcd3 服务发现
    启动时 如果 address 不是 127.0.0.1:2379
    service --registry=etcd --registry_address=127.0.0.1:2379
    
### 使用 protoc 生成代码 
    protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
### 问题列表
