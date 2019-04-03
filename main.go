package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-micro/tool"
)

type platv5AuditConf struct {
	gorm.Model
	Code  string
	Price uint
}

func (m *platv5AuditConf) TableName() string {
	return "platv5_works_13"
}

func main() {

	query := tool.GetMasterConn()

	query.AutoMigrate(&platv5AuditConf{})

	auditConf := platv5AuditConf{}
	query.Model("platv5_works_13").First(&auditConf)

	fmt.Println(auditConf)

	fmt.Println(auditConf.worksId)
}
