package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/auditConf/model"
	auditConf "go-micro/cmd/auditConf/proto"
	"log"
	"time"
)

func main() {
	service := grpc.NewService(
		micro.Name("auditConf"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()

	_ = auditConf.RegisterAuditConfHandler(service.Server(), model.NewAuditConfModel())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
