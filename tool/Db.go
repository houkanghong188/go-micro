package tool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Mysql *gorm.DB

func init() {
	var err error
	Mysql, err = gorm.Open("mysql", "testgroup:testgroupM1@(rm-2zezj9n2lv83nl8x4o.mysql.rds.aliyuncs.com:3306)/makaplatv4?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Mysql.DB().SetMaxIdleConns(3)
	Mysql.DB().SetMaxOpenConns(5)
}
