package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Platv5AuditConf struct {
	ID         int            `gorm:"column:id;primary_key"`
	Name       sql.NullString `gorm:"column:name"`
	Type       sql.NullString `gorm:"column:type"`
	Value      sql.NullString `gorm:"column:value"`
	CreateTime time.Time      `gorm:"column:create_time"`
	UpdateTime time.Time      `gorm:"column:update_time"`
	IsDelete   sql.NullInt64  `gorm:"column:is_delete"`
}

// TableName sets the insert table name for this struct type
func (p *Platv5AuditConf) TableName() string {
	return "platv5_audit_conf"
}
