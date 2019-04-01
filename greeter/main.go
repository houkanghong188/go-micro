package main

import (
	"context"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/greeter/proto"
	"go-micro/tool"
	"log"
	"time"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

type TemplateStoreModel struct {
	TemplateId       string
	DesignerId       int32
	Price            int32
	Version          int32
	Star             int32
	Level            int32
	Sort             int32
	IsVipFree        int32
	TemplateInfo     string
	TemplateProperty string
}

func main() {
	service := grpc.NewService(
		micro.Name("greeter"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	m := &TemplateStoreModel{}
	defer tool.Mysql.Close()
	tool.Mysql.First(m, "id > ? ", 0)

	log.Print(m.TemplateId)

	service.Init()

	_ = greeter.RegisterGreeterHandler(service.Server(), &Greeter{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
