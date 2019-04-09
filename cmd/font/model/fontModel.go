package model

import (
	"context"
	"go-micro/cmd/font/proto/font"
	"go-micro/tool"
)

type FontModel struct {
	Id              int32
	Style           string
	Name            string
	Source          string
	Link            string
	PersonPrice     int32
	EnterPrisePrice int32
	Tax             int32
	Status          int8
	EnableLease     int8
	SaleNumber      int32
	FontIdNo        string
	PreviewTextUrl  string
	PreviewImg      string
	OrderPreviewImg string
}

func (m *FontModel) TableName() string {
	return "platv4_font"
}

func NewFontModel() *FontModel {
	return &FontModel{}
}

func (m *FontModel) Index(ctx context.Context, in *font.IndexRequest, out *font.IndexResponse) error {
	conn := tool.GetMasterConn()

	query := conn.Table(m.TableName())
	if len(in.Ids) > 0 {
		query = query.Where("id in (?)", in.Ids)
	}

	// 必须传 status
	query = query.Where("status = ?", in.GetStatus())

	query.Count(&out.Total)

	var data []*font.Font
	query.Offset(in.Page * in.PageSize).Limit(in.PageSize).Find(&data)
	out.Data = data
	return nil
}
