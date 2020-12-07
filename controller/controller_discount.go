package controller

import (
	"SecKill/model"
	"SecKill/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//首先将数据丢入到redis中，同时异步更新mysql

// @Summary 管理员添加优惠卷
// @Description
// @Tags 优惠卷
// @accept json
// @Produce  json
// @Param object query model.Discount false "请求参数"
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /user/discount [POST]
func CreateDiscount(c *gin.Context){
	var Discount model.Discount
	err := c.ShouldBindJSON(&Discount)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":err,
			"data":nil,
		})
		return
	}
	err = service.CacheDiscount(Discount)
	if err != nil {
		c.JSON(http.StatusBadRequest,"")
		return
	}
	ID := c.GetInt64("Id")
	U := model.UserDiscount{
		UserId:   ID,
		DisCount: Discount,
	}

	//a := service.GetUserIdChanceString(ID)
	service.DISCOUNT <- U
	//b , err := dao.GetSetMembers(a)
	c.JSON(http.StatusOK,"")
}