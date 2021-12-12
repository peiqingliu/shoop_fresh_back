package service

import (
	"errors"
	"fmt"
	"shoop_Fresh_back/config"
	"shoop_Fresh_back/model"
	"shoop_Fresh_back/query"
	"shoop_Fresh_back/repository"
	"shoop_Fresh_back/utils"
	"time"
)


type UserService struct {
	Repo repository.UserRepoInterface
}

type UserSrv interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(user model.User) (*model.User, error)
	ExistByUserID(id string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(u string) (bool,error)
}

func (srv *UserService) List(req *query.ListQuery) (users []*model.User, err error) {
	if req.PageSize < 1 {
		req.PageSize = config.PageSize
	}
	return srv.Repo.List(req)
}

func (srv *UserService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv *UserService) Get(user model.User) (*model.User, error) {
	return srv.Repo.Get(user)
}

func (srv *UserService) ExistByUserID(id string) *model.User {
	return srv.Repo.ExistByUserID(id)
}

func (srv *UserService) Add(user model.User) (*model.User, error) {
	if user.Password == "" {
		user.Password = utils.Md5("123456")
	}
	user.IsDeleted = false
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()
	return srv.Repo.Add(user)
}

func (srv *UserService) Edit(user model.User) (bool, error)  {
	if user.UserId == "" {
		return false, fmt.Errorf("参数错误")
	}
	return srv.Repo.Edit(user)
}

func (srv *UserService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("参数错误")
	}
	user := srv.ExistByUserID(id)
	if user == nil {
		return false, errors.New("参数错误")
	}
	user.IsDeleted = !user.IsDeleted
	return srv.Repo.Delete(*user)
}
