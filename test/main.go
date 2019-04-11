package main

import (
	"context"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/font/proto/font"
	"log"
	"time"
)

func main() {
	service := grpc.NewService(
		micro.Version("1.0.0"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	srv := font.NewFontsService("go.micro.srv.font", service.Client())
	rsp, err := srv.Index(context.Background(), &font.IndexRequest{
		Id: 0,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(len(rsp.Data) == 0)
}

func test_fontService() {

}
