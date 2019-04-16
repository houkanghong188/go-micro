package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-micro/cmd/template/models"
	"go-micro/cmd/template/proto/marketContentConf"
	"go-micro/cmd/template/proto/templateStore"
	"time"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.template"),
		micro.Version("1.0.0"),
		micro.Registry(etcdv3.NewRegistry()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	service.Init()
	// 根据 mark 参数 查询推荐信息， 商城推荐模版列表，归属 模版服务
	marketContentConf.RegisterMarketContentConfHandler(service.Server(), models.NewMarketContentConfModel())
	// 根据 template_ids 查询模版列表
	templateStore.RegisterTemplateStoreServiceHandler(service.Server(), models.NewTemplateStoreModel())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
