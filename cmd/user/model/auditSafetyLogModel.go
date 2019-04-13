package model

import (
	"context"
	"database/sql"
	"go-micro/cmd/user/proto/auditSafetyLog"
	"go-micro/tool"
	"time"
)

type AuditSafetyLogModel struct {
	ID         int            `gorm:"column:id;primary_key"`
	UID        int            `gorm:"column:uid"`
	WorksID    string         `gorm:"column:works_id"`
	CreateTime time.Time      `gorm:"column:create_time"`
	UpdateTime time.Time      `gorm:"column:update_time"`
	IsDelete   int            `gorm:"column:is_delete"`
	Type       sql.NullString `gorm:"column:type"`
	Content    sql.NullString `gorm:"column:content"`
	Scene      sql.NullString `gorm:"column:scene"`
}

func (m AuditSafetyLogModel) TableName() string {
	return "platv5_works_audit_safety_log"
}

func NewAuditSafetyLogModel() *AuditSafetyLogModel {
	return &AuditSafetyLogModel{}
}

func (m *AuditSafetyLogModel) Index(ctx context.Context, req *auditSafetyLog.Request, rsp *auditSafetyLog.Response) error {

	// 获取新的 连接（这里没必要获取，只不过是 举个例子）
	query := tool.GetMasterConn()

	if req.Uid != 0 {
		query = query.Where("uid = ?", req.Uid)
	}

	if req.WorksId != "" {
		query = query.Where("works_id = ?", req.WorksId)
	}

	if req.Scene != "" {
		query = query.Where("scene = ?", req.Scene)
	}

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	query.Table(m.TableName()).Limit(req.Limit).Offset(req.Offset).Find(&rsp.Data)

	return nil
}
