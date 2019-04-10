package model

import (
	"context"
	"errors"
	"go-micro/cmd/works/proto"
	"go-micro/tool"
)

type DailyPvUvModel struct {
	Id         int32
	IsDelete   int32
	Name       string
	Value      string
	Type       string
	CreateTime string
	UpdateTime string
}

func NewDailyPvUvModel() *DailyPvUvModel {
	return &DailyPvUvModel{}
}

func (m *DailyPvUvModel) TableName() string {
	return "platv4_datastory_daily_pvuv_2019"
}

func (m *DailyPvUvModel) Index(ctx context.Context, req *worksAudit.DailyRequest, rsp *worksAudit.DailyIndexResponse) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetStatisticsConn()

	if len(req.WorksIds) == 0 {
		return errors.New("empty rows")
	}

	works := rsp.Data

	query.Table(m.TableName()).Where("event_id in (?)", req.WorksIds).Find(&works)

	rsp.Data = works

	return nil
}

func (m *DailyPvUvModel) Show(ctx context.Context, req *worksAudit.DailyRequest, rsp *worksAudit.DailyResponse) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetStatisticsConn()

	if len(req.WorksId) == 0 {
		return errors.New("empty rows")
	}

	works := rsp.Data

	query.Table(m.TableName()).Where("event_id = ?", req.WorksId).Order("id desc").First(&works)

	rsp.Data = works

	return nil
}
