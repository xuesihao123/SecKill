package controller

import (
	"SecKill/model"
	"SecKill/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//秒杀逻辑
// @Summary 秒杀优惠卷
// @Description 用户对优惠卷进行操作，这里是并发操作
// @Tags 秒杀
// @accept json
// @Produce  json
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /user/seckill/{name} [get]
func SecKillDiscount(c *gin.Context)  {
	//传入数据
	//传入优惠卷的ID
	name := c.Param("name")
	userId := c.GetInt64("Id")
	//从redis中的拿到优惠卷数据
	discount := service.GetDiscount(name)
	userDiscount := model.UserDiscount{
		UserId:   userId,
		DisCount: discount,
	}
	//原子操作去进行redis中优惠卷数据修改
	status , _ := service.SecKillRedis(userDiscount)
	//如果成功就将数据持久化到mysql中，然后返回结果；如果失败就返回结果，这里要注意返回的结果的处理
	if status == 1{
		service.DISCOUNT <- userDiscount
		c.JSON(http.StatusOK , gin.H{
			"status" : true,
			"data" : nil,
			"err" : nil,
		})
		return
	}
	c.JSON(http.StatusBadRequest,gin.H{
		"status" : false,
		"data" : nil,
		"err" : status,
	})
}

