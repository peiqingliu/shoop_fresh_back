package model

import "time"

type User struct {
	UserId string `json:"userId" gorm:"column:user_id"`
	Password string `json:"password" gorm:"column:password"`
	NickName string `json:"nickName" gorm:"column:nick_name"`
	Mobile string `json:"mobile" gorm:"column:mobile" binding:"required"`
	Address string `json:"address",gorm:"column:address"`
	IsDeleted bool `json:"IsDeleted" gorm:"column:is_deleted"`
	CreateUser string `json:"createUser" gorm:"column:create_user"`
	UpdateUser string `json:"updateUser" gorm:"column:update_user"`
	CreateAt  time.Time `json:"createAt" gorm:"column:create_at;default:null"`
	UpdateAt  time.Time `json:"updateAt" gorm:"column:update_at;default:null"`
}