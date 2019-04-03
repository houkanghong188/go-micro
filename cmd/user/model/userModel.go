package model

import (
	"context"
	"errors"
	"go-micro/tool"
)

type UserModel struct {
	Id         int32
	IsDelete   int32
	Name       string
	Value      string
	Type       string
	CreateTime string
	UpdateTime string
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (m *UserModel) TableName() string {
	return "platv5_audit_conf"
}

func (m *UserModel) Show(ctx context.Context, req *User.Request, rsp *User.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()
	query = query.Where("is_delete = 0")

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	data := []*User.Response_Notify{}

	query.Table(m.TableName()).Find(&data)

	rsp.Data = data
	return nil
}
