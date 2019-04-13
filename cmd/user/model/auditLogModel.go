package model

import (
	"context"
	"errors"
	"go-micro/cmd/user/proto/auditLog"
	"go-micro/tool"
)

func NewAuditLogModel() *AuditLogModel {
	return &AuditLogModel{}
}

// TableName sets the insert table name for this struct type

func (m *AuditLogModel) AuditLogIndex(ctx context.Context, req *AuditLog.Request, rsp *AuditLog.AuditLogResponse) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.Uid <= 0 && req.WorksId == "" {
		return errors.New("empty rows")
	}

	if req.Limit == 0 {
		return errors.New("empty rows limit")
	}

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）

	data := []*AuditLog.AuditLogBracket{}

	auditLog := AuditLogModel{}

	if req.Uid != 0 {
		query = query.Where("uid = ?", req.Uid)
	}
	if req.WorksId != "" {
		query = query.Where("works_id = ?", req.WorksId)
	}

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	query.Table(m.TableName()).Count(&rsp.Total)
	query.Table(auditLog.TableName()).Limit(req.Limit).Offset(req.Offset).Find(&data)

	rsp.Data = data
	return nil
}

func (m *AuditLogModel) AuditLogShow(ctx context.Context, req *AuditLog.Request, rsp *AuditLog.LogShowResponse) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.Uid <= 0 && req.WorksId == "" {
		return errors.New("empty rows")
	}

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）

	data := AuditLog.AuditLogBracket{}

	auditLog := AuditLogModel{}

	if req.Uid != 0 {
		query = query.Where("uid = ?", req.Uid)
	}
	if req.WorksId != "" {
		query = query.Where("works_id = ?", req.WorksId)
	}

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	query.Table(auditLog.TableName()).Order("id desc").First(&data)

	rsp.Data = &data
	return nil
}
