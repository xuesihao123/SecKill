package service

import (
	"SecKill/dao"
	"SecKill/model"
	"github.com/jinzhu/gorm"
)

//给用户添加优惠卷
//这里是用户已经抢到优惠卷了
func UserHasDiscount(discount *model.Discount) error {
	if DiscountNameOne(discount.DiscountName){
		Err := model.NewError(4001,"重复添加错误","")
		//添加到日志里
		return Err
	}

	if err := dao.MDb.Create(&discount).Error ; err != nil {
		Err := model.NewError(4000,"添加错误",err.Error())
		//添加到日志里
		return Err
	}
	return nil
}

//给数据库中的优惠卷减一
func DiscountOne(DisCountID int64)	error{
	if err := dao.MDb.Model(&model.Discount{}).Where("discount_id",DisCountID).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1)).Error ; err != nil {
		Err := model.NewError(4000,"修改错误",err.Error())
		//添加到日志里
		return Err
	}
	return nil
}

func DiscountNameOne(Name string) bool{
	if err := dao.MDb.Where("discount_name = ?", Name).First(&model.Discount{}).Error ; err != nil{
		return false
	}
	return true
}