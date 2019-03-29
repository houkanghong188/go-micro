package main

import (
	"context"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	greeter "go-micro/greeter/proto"
	"log"
	"time"
)


type Greeter struct {

}

func (g *Greeter) Hello(ctx context.Context, req  *greeter.Request, rsp *greeter.Response) error  {
	rsp.Msg = "Hello " + req.Name
	return nil
}



func main() {
	service := grpc.NewService(
		micro.Name("greeter"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30 * time.Second),
		micro.RegisterInterval(10 * time.Second),
	)

	service.Init()

	_ = greeter.RegisterGreeterHandler(service.Server(), &Greeter{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}