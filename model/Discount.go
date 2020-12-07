package model

import "SecKill/dao"

//取出所有的优惠卷
func FindDiscount()[]Discount {
	var discounts []Discount
	if err := dao.MDb.Find(&discounts).Error;err != nil{
		panic("数据预加载失败")
	}
	return discounts
}
