package dao

import (
	"SecKill/config"
	"github.com/go-redis/redis"
)

var (
	RDb *redis.Client
)

func RedisInit() error{
	RedisString := config.RedisString()
	RDb = redis.NewClient(&redis.Options{
		Addr:     RedisString,
		Password: config.RedisDao.Password,
		DB:       0,
	})
	_ ,err := RDb.Ping().Result()
	if err != nil {
		panic("lk")
	}
	return nil
}


//给Hash表中添加数据
func SetMapForever(key string, field map[string]interface{}) (string, error) {
	return RDb.HMSet(key, field).Result()
}

//从Hash表中取数据
func GetMap(key string, fields ...string) ([]interface{}, error) {
	return RDb.HMGet(key, fields...).Result()
}

//给redis的set添加数据
func SetAdd(key string, field string) (int64, error) {
	return RDb.SAdd(key, field).Result()
}

func SetIsMember(key string, field string) (bool, error) {
	return RDb.SIsMember(key, field).Result()
}

func GetSetMembers(key string) ([]string, error) {
	return RDb.SMembers(key).Result()
}

func uniqueOccurrences(arr []int) bool {
	num := make(map[int]int)
	simple := make(map[int]int)
	for _ , value := range arr{
		num[value] = num[value]+1
	}

	for _ ,value := range num{
		simple[value] = simple[value]+1
		if simple[value] > 1{
			return false
		}
	}
	return true
}

