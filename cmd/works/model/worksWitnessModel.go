package model

import (
	"context"
	"errors"
	witness "go-micro/cmd/works/proto/witness"
	"go-micro/tool"
)

// 作品举报表
type WorksWitnessesModel struct {
	Id         int32
	IsDelete   int32
	Name       string
	Value      string
	Type       string
	CreateTime string
	UpdateTime string
}

// TableName sets the insert table name for this struct type
func (m *WorksWitnessesModel) TableName() string {
	return "platv4_witnesses"
}

func NewWorksWitnessesModel() *WorksWitnessesModel {
	return &WorksWitnessesModel{}
}

func (m *WorksWitnessesModel) Index(ctx context.Context, req *witness.WitnessesRequest, rsp *witness.WitnessesResponse) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	query.Table(m.TableName()).Count(&rsp.Total)

	if req.Limit == 0 {
		return errors.New("empty rows")
	}

	var data []*witness.Witness

	query = query.Table(m.TableName())

	if req.Uid >= 0 {
		query = query.Where("uid = ?", req.Uid)
	}

	if len(req.WorksId) > 0 {
		query = query.Where("event_id = ?", req.WorksId)
	}

	query = query.Limit(req.Limit).Offset(req.Offset)

	query.Count(rsp.Total)

	query.Find(&data)

	rsp.Data = data

	return nil
}
