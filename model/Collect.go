package model

import "SecKill/dao"

//添加一个优惠卷关注
func AddCollectOne(collect *Collect)error{
	if err := dao.MDb.Create(&collect).Error ; err != nil {
		Err := NewError(4000 , "添加失败",err.Error())
		return Err
	}
	return nil
}

//取消关注
func DeleteCollect(ID int64) error{
	if err := dao.MDb.Where("collect_id = ?",ID).Delete(&Collect{}).Error ; err != nil {
		Err := NewError(4000 , "删除失败",err.Error())
		return Err
	}
	return nil
}
