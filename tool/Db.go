package tool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-micro/config"
)

var MysqlMaster *gorm.DB
var MysqlSlaver *gorm.DB
var MysqlStatistics *gorm.DB

func init() {
	var err error
	// 主库
	MysqlMaster, err = gorm.Open("mysql", config.MysqlMasterDns)
	if err != nil {
		panic(err)
	}
	MysqlMaster.DB().SetMaxIdleConns(config.MysqlMasterMaxIdleConns)
	MysqlMaster.DB().SetMaxOpenConns(config.MysqlMasterMaxOpenConns)

	// 从库
	MysqlSlaver, err = gorm.Open("mysql", config.MysqlSlaverDns)
	if err != nil {
		panic(err)
	}

	MysqlSlaver.DB().SetMaxIdleConns(config.MysqlSlaverMaxIdleConns)
	MysqlSlaver.DB().SetMaxOpenConns(config.MysqlSlaverMaxOpenConns)

	// 统计库
	MysqlStatistics, err = gorm.Open("mysql", config.MysqlStatisticsDns)
	if err != nil {
		panic(err)
	}
	MysqlStatistics.DB().SetMaxIdleConns(config.MysqlMasterMaxIdleConns)
	MysqlStatistics.DB().SetMaxOpenConns(config.MysqlMasterMaxOpenConns)
}

func GetMasterConn() *gorm.DB {
	return MysqlMaster
}

func GetStatisticsConn() *gorm.DB {
	return MysqlStatistics
}

func GetSlavelConn() *gorm.DB {
	return MysqlSlaver
}
