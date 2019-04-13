package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/template/models"
	"go-micro/cmd/template/proto/marketContentConf"
	"time"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.marketContentConf"),
		micro.Version("1.0.0"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()
	marketContentConf.RegisterMarketContentConfHandler(service.Server(), models.NewMarketContentConfModel())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
