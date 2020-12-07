package service

import (
	"SecKill/dao"
	"SecKill/model"
	"fmt"
	"strconv"
)

func CacheDiscount(discount model.Discount)(err error){
	if _ ,err := CacheDiscountId(discount) ; err != nil{
		panic(err.Error())
		return err
	}
	_ , err = CacheDiscountDate(discount)
	if err != nil {
		return err
	}
	return nil
}

func CacheDiscountId(discount model.Discount)(int64, error){
	//将数据加入hsah中不加入set中
	UserIdString := GetUserIdChanceString(discount.DiscountUserID)
	DiscountString := GetDiscountChanceString(discount.DiscountName)
	if val , err := dao.SetAdd(UserIdString,DiscountString) ; err != nil {
		//加入日志
		return val,err
	}
	return  0, nil
}

func CacheDiscountDate(discount model.Discount)(string , error) {
	fields := map[string]interface{}{
			"DiscountId":        discount.DiscountId,
			"DiscountName":      discount.DiscountName,
			"DiscountStatus":    discount.DiscountStatus,
			"DiscountStartTime": discount.DiscountStartTime,
			"DiscountEndTime":   discount.DiscountEndTime,
			"DiscountNum":       discount.DiscountNum,
			"DiscountUserID":    discount.DiscountUserID,
			"DiscountStock":     discount.DiscountStock,
			"DiscountContext":   discount.DiscountContext,
	}
	DiscountString := GetDiscountChanceString(discount.DiscountName)
	val, err := dao.SetMapForever(DiscountString, fields)
	return val, err
}

func GetDiscountChanceString(name string)string{
	return fmt.Sprintf("Discount-%s", name)
}

func GetUserIdChanceString(Id int64)string{
	return fmt.Sprintf("User-%d", Id)
}

func GetDiscount(name string)model.Discount{
	name = GetDiscountChanceString(name)
	values , _ := dao.GetMap(name,	"DiscountId", "DiscountName", "DiscountStatus", "DiscountStartTime", "DiscountEndTime", "DiscountNum", "DiscountUserID", "DiscountStock", "DiscountContext")
	ID, _ := strconv.ParseInt(values[0].(string), 10, 64)
	Status, _ := strconv.ParseInt(values[2].(string), 10, 64)
	Num, _ := strconv.ParseInt(values[5].(string), 10, 64)
	UserID, _ := strconv.ParseInt(values[6].(string), 10, 64)
	Stock, _ := strconv.ParseInt(values[7].(string), 10, 64)

	return model.Discount{
		DiscountId:        ID,
		DiscountName:      values[1].(string),
		DiscountStatus:    Status,
		DiscountStartTime: values[3].(string),
		DiscountEndTime:   values[4].(string),
		DiscountNum:       Num,
		DiscountUserID:    UserID,
		DiscountStock:     Stock,
		DiscountContext:   values[8].(string),
	}
}

//使用lua对redis进行原子操作
func SecKillRedis(discount model.UserDiscount)(int64 , error){
	UserIdString := GetUserIdChanceString(discount.UserId)
	DiscountString := GetDiscountChanceString(discount.DisCount.DiscountName)
	res , err := EvalSHA(SecKillSHA,[]string{UserIdString,DiscountString})
	if err != nil {
		return 0 , nil
	}
	val := res.(int64)
	return val, nil
}