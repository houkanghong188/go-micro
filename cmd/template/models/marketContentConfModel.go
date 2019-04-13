package models

import (
	"context"
	"fmt"
	"go-micro/cmd/template/proto/marketContentConf"
	"go-micro/tool"
)

/**
`id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(32) DEFAULT NULL COMMENT '标签名称',
  `type` enum('h5','video','poster','text') NOT NULL DEFAULT 'text' COMMENT '类型',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `sort` smallint(6) DEFAULT NULL COMMENT '排序',
  `parent_id` int(11) DEFAULT NULL COMMENT '父类id',
  `mark` enum('pc_comprehensive','app_comprehensive','pc_category','app_category','client_comprehensive','client_category') NOT NULL COMMENT '标示',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `is_delete` smallint(6) DEFAULT '0' COMMENT '是否删除',
  `relevant_category` int(11) NOT NULL COMMENT '关联分类信息',
*/
type MarketContentConfModel struct {
	Id               int32
	Title            string
	Type             int
	Description      string
	Sort             int16
	ParentId         int32
	Mark             int
	CreateTime       string
	UpdateTime       string
	IsDelete         int8
	RelevantCategory int32
	Recommend        MarketContentRecommendModel
}

type UserModel struct {
}

func NewMarketContentConfModel() *MarketContentConfModel {
	return &MarketContentConfModel{}
}

func (m *MarketContentConfModel) TableName() string {
	return "platv5_market_content_conf"
}

func (m *MarketContentConfModel) Index(ctx context.Context, in *marketContentConf.Request, out *marketContentConf.Response) error {
	conn := tool.GetMasterConn()

	conn.Table(m.TableName()).Where("mark in (?) and status = 1", in.Mark).Find(&out.ContentConf)

	var content []*marketContentConf.ContentConf
	var recommend2 []*marketContentConf.ContentRecommend
	conn.Table(m.TableName()).Where("mark in (?) and status = 1", in.Mark).Related(&recommend2).Find(&out.ContentConf)
	conn.Model(&content).Related(&recommend2)
	fmt.Println(recommend2)

	var ids []int32
	for _, v := range out.ContentConf {
		ids = append(ids, v.Id)
	}

	fmt.Println(ids)

	//conn = tool.GetSlavelConn()
	////m2 := NewMarketContentRecommendModel()
	//
	//var recommend []*marketContentConf.ContentRecommend
	////conn.Table(m2.TableName()).Where("content_id in (?) and is_delete = 0", ids).Find(&recommend)
	//
	//reMap := make(map[int32][]*marketContentConf.ContentRecommend)
	//for _,v := range recommend{
	//	reMap[v.ContentId] = append(reMap[v.ContentId],v)
	//}
	//
	//for k,v := range out.ContentConf {
	//	if res, ok := reMap[v.Id]; ok {
	//		out.ContentConf[k].Recommend = res
	//	}
	//}

	return nil
}
