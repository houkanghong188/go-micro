package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/Works/proto"
	model2 "go-micro/cmd/auditConf/model"
	"go-micro/cmd/auditConf/proto"
	"go-micro/cmd/works/model"
	"log"
	"time"
)

func main() {

	service := grpc.NewService(
		micro.Name("go.micro.srv.worksAudit"),
		micro.Version("1.0.0"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()

	_ = worksAudit.RegisterWorksAuditHandler(service.Server(), model.NewWorksAuditModel())
	_ = worksAudit.RegisterWorksHandler(service.Server(), model.NewWorksModel())
	_ = worksAudit.RegisterDailyPvUvHandler(service.Server(), model.NewDailyPvUvModel())
	// 临时兼容注册到一个服务中
	_ = auditConf.RegisterAuditConfHandler(service.Server(), model2.NewAuditConfModel())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
