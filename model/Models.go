package model

import "SecKill/dao"

const (
	CommonUser = 1
	AdminUser = 2
	DeleteUser = 3
	FailUser = 4

	Wait = 1
	Activity = 2
	End = 3
)

// User 用户的model
type User struct {
	UserId int64 `gorm:"primary_key" json:"user_id"`
	UserName string `gorm:"type:varchar(20)" json:"user_name"`
	UserPassword string `gorm:"type:varchar(100)" json:"user_password"`
	UserStatus int64 ` json:"user_status"`
}

// Discount 优惠卷的model
type Discount struct {
	DiscountId int64 `gorm:"primary_key" json:"discount_id"`
	DiscountName string `gorm:"type:varchar(20)" json:"discount_name"`						//优惠卷名
	DiscountStatus int64 ` json:"discount_status"`						//状态
	DiscountStartTime string `gorm:"type:dateTime" json:"discount_start_time"`				//开启时间
	DiscountEndTime string `gorm:"type:dateTime" json:"discount_end_time"`					//结束时间
	DiscountNum int64 ` json:"discount_num"`
	DiscountUserID int64 ` json:"discount_user_id"`						//拥有者
	DiscountStock   int64 `json:"discount_stock"`       									//面额
	DiscountContext string `gorm:"type:text" json:"discount_context"`						//内容
}

// Collect 收藏列表
type Collect struct {
	CollectId int64 `gorm:"primary_key" json:"collect_id"`
	CollectUserId int64 `gorm:"index" json:"collect_user_id"`
	CollectDiscountId int64 `gorm:"index" json:"collect_discount_id"`
}

func DateInit(){
	//映射结构体
	dao.MDb.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(User{})
	dao.MDb.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(Discount{})
	dao.MDb.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(Collect{})
	//添加外键
	dao.MDb.Model(&Discount{}).AddForeignKey("discount_user_id","user(user_id)","RESTRICT","RESTRICT")
	dao.MDb.Model(&Collect{}).AddForeignKey("collect_user_id","user(user_id)","RESTRICT","RESTRICT")
	dao.MDb.Model(&Collect{}).AddForeignKey("collect_discount_id","discount(discount_id)","RESTRICT","RESTRICT")
}