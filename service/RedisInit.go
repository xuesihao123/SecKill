package service

import (
	"SecKill/dao"
	"SecKill/model"
)
var SecKillSHA string

const SeckillScicpt = `
	--KEY[1] 为用户的IdSTRING
	--KEY[2] 为优惠卷的NameSTRING

	local	NUM = redis.call("hget" , KEYS[2] , "DiscountNum")
	if(NUM == false)
	then
		return -1
		--没有这个优惠卷
	end
	
	if(tonumber(NUM) == 0)
	then
		return -2
		--优惠卷已经抢完
	end
	
	local userHasCoupon = redis.call("SISMEMBER", KEYS[1], KEYS[2]);
	if(userHasCoupon == 1)
	then
		return -3
		--已经有这个优惠卷
	end
	
	redis.call("hset", KEYS[2], "DiscountNum", NUM - 1);
	redis.call("SADD", KEYS[1], KEYS[2]);
	return 1;
`
const MAX = 10000
var DISCOUNT = make(chan model.UserDiscount , MAX)
//需要将Mysql中优惠卷的数据预加载进入redis
//分别加载一个Map[string]string，一个Map[string]struct

func ScriptInit()  {
	SecKillSHA = PrepareScript(SeckillScicpt)
	RedisInitStruct()
}

func RedisInitStruct(){
	discounts := model.FindDiscount()
	for _ , discount := range discounts{
		err := CacheDiscount(discount)
		if err != nil {
			panic("预加载出错")
		}
	}
}

//设置api去实现数据库内容更新
//这里操作数据库要开启事务对mysql进行更新
//首先更新redis中的数据
func UpdateRedisToMysql(){
	for{
		discount := <-DISCOUNT
		discount.DisCount.DiscountNum  = 1
		discount.DisCount.DiscountUserID = discount.UserId
		discount.DisCount.DiscountStatus = model.CommonUser
		tx := dao.MDb.Begin()
		if err := UserHasDiscount(&discount.DisCount) ; err != nil{
			tx.Rollback()
			//日志
			return
		}
		if err:= DiscountOne(discount.DisCount.DiscountUserID) ; err != nil{
			tx.Rollback()
			//日志
			return
		}
		return
	}

}

var flag = false
//开启消费者
func RunUpdate()  {
	go UpdateRedisToMysql()
	flag = true
}

//加载LUA程序
func PrepareScript(script string) string {
	// sha := sha1.Sum([]byte(script))
	scriptsExists, err := dao.RDb.ScriptExists(script).Result()

	if err != nil {
		panic("Failed to check if script exists: " + err.Error())
	}
	if !scriptsExists[0] {
		scriptSHA, err := dao.RDb.ScriptLoad(script).Result()
		if err != nil {
			panic("Failed to load script " + script + " err: " + err.Error())
		}
		return scriptSHA
	}
	print("Script Exists.")
	return ""
}

func EvalSHA(sha string, args []string) (interface{}, error) {
	val, err := dao.RDb.EvalSha(sha, args).Result()
	if err != nil {
		print("Error executing evalSHA... " + err.Error())
		return nil, err
	}
	return val, nil
}