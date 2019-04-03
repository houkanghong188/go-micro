package main

import (
	"github.com/jinzhu/gorm"
)

type platv5AuditConf struct {
	gorm.Model
	Code  string
	Price uint
}

func (m *platv5AuditConf) TableName() string {
	return "platv5_works_13"
}
