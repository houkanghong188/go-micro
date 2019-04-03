package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-micro/tool"
	"strconv"
	"time"
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
	return "platv5_works_audit"
}

func (m *DailyPvUvModel) Show(ctx context.Context, req *DailyPvUv.Request, rsp *DailyPvUv.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	data := []*DailyPvUv.Response_Notify{}

	query.Table(m.TableName()).Find(&data)

	for k, v := range data {
		tableName := "platv5_works_" + strconv.Itoa(int(v.Uid%16))
		work := works{}
		query.Table(tableName).Where("works_id = ?", v.EventId).First(&work)

		fmt.Println(work)

		fmt.Println(work == (works{}))
		if work == (works{}) {
			fmt.Println(work)
			data[k].Content = " "
			data[k].Title = " "
		} else {
			data[k].Title = work.Title
			data[k].Content = work.Content
		}
	}

	rsp.Data = data

	return nil
}

func (m *DailyPvUvModel) Index(ctx context.Context, req *DailyPvUv.Request, rsp *DailyPvUv.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	query.Table(m.TableName()).Count(&rsp.Total)

	if req.PageSize == 0 {
		return errors.New("empty rows")
	}

	data := []*DailyPvUv.Response_Notify{}

	query.Table(m.TableName()).Limit(req.PageSize).Offset(req.PageSize * req.Page).Find(&data)

	for k, v := range data {
		tableName := "platv5_works_" + strconv.Itoa(int(v.Uid%16))
		work := works{}
		query.Table(tableName).Where("works_id = ?", v.EventId).First(&work)

		fmt.Println(work)

		fmt.Println(work == (works{}))
		if work == (works{}) {
			fmt.Println(work)
			data[k].Content = " "
			data[k].Title = " "
		} else {
			data[k].Title = work.Title
			data[k].Content = work.Content
		}
	}

	rsp.Data = data

	return nil
}
