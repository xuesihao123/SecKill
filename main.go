package main

import (
	"SecKill/dao"
	"SecKill/model"
	"SecKill/router"
	"SecKill/service"
)

func main() {
	dao.MysqlInit()
	dao.RedisInit()
	model.DateInit()
	service.ScriptInit()
	r := router.RoutersInit()

	r.Run(":9111")
}
