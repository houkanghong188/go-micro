package main

import (
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/grpc"
	_ "github.com/micro/go-plugins/client/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	_ "github.com/micro/go-plugins/server/grpc"
	_ "github.com/micro/go-plugins/transport/grpc"
	"github.com/micro/micro/cmd"
	"time"
)

func main() {

	cmd.Init(
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterInterval(10 * time.Second),
		micro.RegisterTTL(30 * time.Second),
		)
}
