package model

import (
	"context"
	"errors"
	"go-micro/cmd/auditConf/proto"
	"go-micro/tool"
)

type AuditConfModel struct {
	Id         int32
	IsDelete   int32
	Name       string
	Value      string
	Type       string
	CreateTime string
	UpdateTime string
}

func NewAuditConfModel() *AuditConfModel {
	return &AuditConfModel{}
}

func (m *AuditConfModel) TableName() string {
	return "platv5_audit_conf"
}

func (m *AuditConfModel) Show(ctx context.Context, req *auditConf.Request, rsp *auditConf.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()
	query = query.Where("is_delete = 0")

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	data := []*auditConf.Response_Notify{}

	query.Table(m.TableName()).Find(&data)

	rsp.Data = data
	return nil
}

func (m *AuditConfModel) Create(ctx context.Context, req *auditConf.Request, rsp *auditConf.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	Mysql := tool.GetMasterConn()

	auditConf := AuditConfModel{Name: req.Name, Type: req.Type, Value: req.Value}

	Mysql.Create(&auditConf)

	return nil
}

func (m *AuditConfModel) Update(ctx context.Context, req *auditConf.Request, rsp *auditConf.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.Id <= 0 {
		return errors.New("empty rows")
	}

	var auditConf = AuditConfModel{}
	query.Where("id = ?", req.Id).First(&auditConf)

	query.Model(auditConf).Updates(req)

	return nil
}

func (m *AuditConfModel) Index(ctx context.Context, req *auditConf.Request, rsp *auditConf.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	query.Table(m.TableName()).Count(&rsp.Total)

	if req.PageSize == 0 {
		return errors.New("empty rows")
	}

	data := []*auditConf.Response_Notify{}

	query.Table(m.TableName()).Limit(req.PageSize).Offset(req.PageSize * req.Page).Find(&data)

	rsp.Data = data

	return nil
}
