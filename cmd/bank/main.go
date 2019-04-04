package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/bank/model"
	"go-micro/cmd/bank/proto"
	"log"
	"time"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.bank"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()

	// 可直接使用 model 实现 proto interface
	_ = bank.RegisterBankHandler(service.Server(), model.NewBankModel())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
