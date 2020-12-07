package model

// RUser 登录响应数据
type RUser struct {
	UserId int64 `json:"user_id"`
	UserName string `json:"user_name"`
	UserStatus int64 `json:"user_status"`
}

// LUser 登录请求数据
type LUser struct {
	UserName string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

// RDiscount 登录响应数据
type RDiscount struct {
	DiscountId int64 `json:"discount_id"`
	DiscountName string `json:"discount_name"`
	DiscountStatus int64 `json:"discount_status"`
	DiscountStartTime string `json:"discount_start_time"`
	DiscountEndTime string `gorm:"type:dateTime" json:"discount_end_time"`
	DiscountNum int64 `json:"discount_num"`
	DiscountUserID int64 `json:"discount_user_id"`
	DiscountUserName string `json:"discount_user_name"`
	DiscountStock   int64 `json:"discount_stock"`
	DiscountContext string `json:"discount_context"`
}

//返回用户拥有优惠卷
type HasDiscount struct {
	DiscountId int64 `json:"discount_id"`
	DiscountName string `json:"discount_name"`
	DiscountStatus int64 `json:"discount_status"`
	DiscountEndTime string `gorm:"type:dateTime" json:"discount_end_time"`
	DiscountStock   int64 `json:"discount_stock"`
	DiscountUserID int64 `json:"discount_user_id"`
	DiscountUserName string `json:"discount_user_name"`
	DiscountContext string `json:"discount_context"`
}

type UserDiscount struct {
	UserId int64
	DisCount Discount
}