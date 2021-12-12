package model

import "time"

type Product struct {
	ProductId string `json:"productId" gorm:"column:product_id"`
	ProductName string `json:"productName" gorm:"column:product_name"`
	ProductIntro string `json:"ProductIntro" gorm:"column:product_intro"`
	CategoryId string `json:"categoryId" gorm:"column:category_id"`
	OriginalPrice string `json:"originalPrice" gorm:"column:original_price"`
	SellingPrice string `json:"SellingPrice" gorm:"column:selling_price"`
	CreateUser string `json:"createUser" gorm:"column:create_user"`
	CreateAt time.Time `json:"CreateAt" gorm:"column:create_at"`
	UpdateUser string `json:"updateUser" gorm:"column:update_user"`
	UpdateAt time.Time `json:"UpdateAt" gorm:"column:update_at"`
	IsDeleted string `json:"isDeleted" gorm:"column:is_deleted"`
}
