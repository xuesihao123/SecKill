package controller

import (
	"SecKill/model"
	routers "SecKill/router/middleware"
	"SecKill/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 登录接口
// @Description 一个拉胯的登录接口（后台提供额外一个正则表达式匹配）
// @Tags 用户操作
// @accept json
// @Produce  json
// @Param object query model.LUser false "请求参数"
// @Success 200 {object} model.RUser {"status":true,"data":"数据",err:null}
// @Failure 400 {object} model.RUser {"status":false,"data":null,err:"错误数据"}
// @Router /user/main [POST]
func Login(c *gin.Context){
	var User model.LUser
	err := c.ShouldBindJSON(&User)//获取登陆时的请求参数
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":err,
			"data":nil,
		})
		return
	}
	User.UserPassword = service.MD5INIT(User.UserPassword)
	Ruser , Err := model.SelectLogin(&User)
	if Err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":Err,
			"data":nil,
		})
		return
	}
	token , _ := routers.GenToken(Ruser)
	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"err": nil,
		"Token":token,
		"data": Ruser,
	})
}

// @Summary 注册接口
// @Description 一个拉胯的注册接口（后台提供额外一个正则表达式匹配）
// @Tags 用户操作
// @accept json
// @Produce  json
// @Param object query model.User false "请求参数"
// @Success 200 {object}  {"status":true,"data":null,err:null}
// @Failure 400 {object}  {"status":false,"data":null,err:"错误数据"}
// @Router /user/register [POST]
func Register(c *gin.Context){
	var User model.User
	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":err,
			"data":nil,
		})
		return
	}
	User.UserPassword = service.MD5INIT(User.UserPassword)
	Err := model.CreateUser(&User)
	if Err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":nil,
			"data":nil,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"err": nil,
		"data": nil,
	})
}

// @Summary 冻结接口
// @Description 一个拉胯的冻结接口
// @Tags 用户操作
// @accept json
// @Produce  json
// @Param object query model.User false "请求参数"
// @Success 200 {object}  {"status":true,"data":null,err:null}
// @Failure 400 {object}  {"status":false,"data":null,err:"错误数据"}
// @Router /user/fail [POST]
func FailUser(c *gin.Context){
	var User model.RUser
	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":err,
			"data":nil,
		})
		return
	}
	Err := model.UpdateUserStatus(User.UserId)
	if Err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":false,
			"err":nil,
			"data":nil,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"err": nil,
		"data": nil,
	})
}

