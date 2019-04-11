package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/font/model"
	"go-micro/cmd/font/proto/font"
	"log"
	"time"
)

func main() {

	service := grpc.NewService(
		micro.Name("go.micro.srv.font"),
		micro.Version("1.0.0"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()
	_ = font.RegisterFontsHandler(service.Server(), model.NewFontModel())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
