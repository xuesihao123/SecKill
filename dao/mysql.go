package dao

import (
	"SecKill/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	MDb *gorm.DB
)

func MysqlInit()(err error){
	MysqlString := config.MysqlString()
	MDb,err = gorm.Open("mysql",MysqlString)
	if err != nil {
		//添加日志
		panic(err)
		return
	}
	//全局禁用表名创建为复数
	MDb.SingularTable(true)
	return MDb.DB().Ping()
}
