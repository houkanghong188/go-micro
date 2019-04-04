package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/auditConf/model"
	"go-micro/cmd/auditConf/proto"
	model2 "go-micro/cmd/bank/model"
	"go-micro/cmd/bank/proto"
	"log"
	"time"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.auditConf"),
		micro.Version("1.0.0"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()

	_ = auditConf.RegisterAuditConfHandler(service.Server(), model.NewAuditConfModel())
	bank.RegisterBankHandler(service.Server(), model2.NewBankModel())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
