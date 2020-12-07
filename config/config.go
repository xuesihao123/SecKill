package config

//数据库参数
type Dao struct{
	Username string
	Password string
	Port string
	Host string
}

var (
	MysqlDao = Dao{
		Username: "root",
		Password: "123456",
		Port:     "3307",
		Host: "127.0.0.1" ,
	}
	MysqlDataBase = "seckill"
	RedisDao = Dao{
		Username: "",
		Password: "",
		Port:     "6379",
		Host:     "localhost",
	}
)

func MysqlString()string{
	mysql := MysqlDao.Username + ":" + MysqlDao.Password + "@(" +MysqlDao.Host + ":" + MysqlDao.Port + ")/" + MysqlDataBase + "?charset=utf8&parseTime=True&loc=Local"
		return mysql
}

func RedisString()string{
	redisAddr := RedisDao.Host + ":" + RedisDao.Port
		return redisAddr
}
