package model

import "SecKill/dao"

//登陆
func SelectLogin(UserL *LUser)(*RUser , error){
	var user User
	var U RUser
	if SelectLife(UserL.UserName) == false{
		Err := NewError(4002,"账号不存在","")
		return nil ,Err
	}
	if err := dao.MDb.Where("user_name = ? and user_password = ?" , UserL.UserName,UserL.UserPassword).First(&user).Error; err != nil {
		Err := NewError(4001,"账号或密码错误",err.Error())
		return nil,Err
	}
	U.UserName = user.UserName
	U.UserId = user.UserId
	U.UserStatus = user.UserStatus
	return &U ,nil
}

//查询用户是否存在
func SelectLife(UserName string)bool{
	if err := dao.MDb.Where("user_name = ?",UserName).First(&User{}).Error; err != nil {
		return false
	}
	return true
}

//注册
func CreateUser(user *User)error{
	if SelectLife(user.UserName){
		Err := NewError(4003,"用户已经存在","")
		return Err
	}
	if err := dao.MDb.Create(user).Error ; err!=nil{	
		Err := NewError(4000,"未知错误",err.Error())
		return Err
	}
	return nil
}

//冻结用户
func UpdateUserStatus(id int64) error {
	if err:=dao.MDb.Model(&User{}).Where("user_id = ?",id).Update("user_status",FailUser).Error ;err!=nil{
		Err := NewError(4000,"未知错误",err.Error())
		return Err
	}
	return nil
}

func DeleteUserStatus(id int64) error {
	if err:=dao.MDb.Model(&User{}).Where("user_id = ?",id).Update("user_status",DeleteUser).Error ;err!=nil{
		Err := NewError(4000,"未知错误",err.Error())
		return Err
	}
	return nil
}

func SelectUserID(ID int64)(user *User , err error){
	if err = dao.MDb.Where("user_id" , ID).Error ; err != nil{
		Err := NewError(4000, "没有用户",err.Error())
		return nil , Err
	}
	return user , nil
}

func IsAdmin(status int64)bool{
	if status == AdminUser{
		return true
	}
	return false
}