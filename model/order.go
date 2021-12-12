package model

import "time"

type Order struct {
	OrderId string `json:"orderId gorm:"column:order_id"`
	UserId string `json:"userId" gorm:"column:user_id"`
	NickName string `json:"nickName" gorm:"column:nick_name"`
	Mobile string `json:"mobile" gorm:"column:mobile"`
	TotalPrice string `json:"totalPrice" gorm:"column:total_price"`
	PayStatus string `json:"payStatus" gorm:"column:pay_status"`
	PayType string `json:"payType" gorm:"column:pay_type"`
	PayTime time.Time `json:"payTime" gorm:"column:pay_time"`
	UserAddress string `json:"userAddress" gorm:"column:user_address"`
	IsDeleted string `json:"isDeleted" gorm:"column:is_deleted"`
}
