package repository

import (
	"fmt"
	"gorm.io/gorm"
	"shoop_Fresh_back/model"
	"shoop_Fresh_back/query"
	"shoop_Fresh_back/utils"
)

type UserRepository struct {
	DB *gorm.DB
}

// 接口
type UserRepoInterface interface {
	List(req *query.ListQuery) (users [] *model.User, err error)
	GetTotal(req *query.ListQuery) (total int,err error)
	ExistByUserID(id string) *model.User
	Get(user model.User) (*model.User, error)
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(u model.User) (bool,error)
}

func (repo *UserRepository) List(req *query.ListQuery) (users [] *model.User, err error) {
	fmt.Println(req)
	db := repo.DB
	limit,offset := utils.Page(req.PageSize,req.PageNum) // 分页
	if 	err := db.Order("user_id desc").Limit(limit).Offset(offset).Find(&users).Error;err != nil {
		return nil, err
	}
	return users,nil
}

func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int,err error)  {
	fmt.Println(req)
	var users []model.User
	var total1 int64
	db:=repo.DB
	if err := db.Find(&users).Count(&total1).Error;err != nil {
		return int(total1), err
	}
	return total, nil
}


func (repo *UserRepository) ExistByUserID(id string) *model.User {
	var user model.User
	repo.DB.Where("user_id = ?", id).First(&user)
	return &user
}

func (repo *UserRepository) Get(user model.User) (*model.User, error)  {
	if err:=repo.DB.Where(&user).Find(&user).Error;err != nil {
		return nil,err
	}
	return &user,nil
}

func (repo *UserRepository) Add(user model.User) (*model.User, error)  {
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil,fmt.Errorf("用户添加失败")
	}
	return &user,nil
}

func (repo *UserRepository) Edit(user model.User) (bool, error)  {
	err := repo.DB.Model(&user).Updates(map[string]interface{}{"nick_name": user.NickName, "mobile": user.Mobile, "address": user.Address}).Error
	if err != nil{
		return false, err
	}
	return true, err
}

func (repo *UserRepository) Delete(user model.User) (bool,error) {
	err := repo.DB.Model(&user).Where("user_id = ?", user.UserId).Update("IsDeleted", user.IsDeleted).Error
	if err != nil{
		return false, err
	}
	return true, err
}