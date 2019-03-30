package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	bank "go-micro/cmd/bank/proto"
	"log"
	"time"
	"go-micro/cmd/bank/model"
)


func main() {
	service := grpc.NewService(
		micro.Name("bank"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30 * time.Second),
		micro.RegisterInterval(10 * time.Second),
	)

	service.Init()

	_ = bank.RegisterBankHandler(service.Server(), model.NewBankModel())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}