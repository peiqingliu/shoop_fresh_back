package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"shoop_Fresh_back/enum"
	"shoop_Fresh_back/model"
	"shoop_Fresh_back/query"
	"shoop_Fresh_back/resp"
	"shoop_Fresh_back/service"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (h *UserHandler) GetEntity(result model.User) resp.UserResp{
	return resp.UserResp{
		Id:        result.UserId,
		Key:       result.UserId,
		UserId:    result.UserId,
		NickName:  result.NickName,
		Mobile:    result.Mobile,
		Address:   result.Address,
		IsDeleted: result.IsDeleted,
	}
}

// 分页查询
func (h *UserHandler) UserListHandler(c *gin.Context) {
	var q query.ListQuery
	// 默认返回失败
	result := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	// 绑定参数
	err := c.ShouldBind(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"result": result})
	}
	list, err :=h.UserSrv.List(&q)
	total, err := h.UserSrv.GetTotal(&q)
	if err != nil {
		logrus.Error(err)
	}
	if q.PageSize == 0 {
		q.PageSize = 10
	}
	ret := int(total % q.PageSize)
	ret2 := int(total / q.PageSize)
	totalPage := 0
	if ret == 0 {
		totalPage = ret2
	}else {
		totalPage = ret2 + 1
	}
	var newList []*resp.UserResp
	for _, item := range list  {
		r := h.GetEntity(*item)
		newList = append(newList, &r)
	}
	result = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     total,
		TotalPage: totalPage,
		Data:      newList,
	}
	c.JSON(http.StatusOK,gin.H{"entity": result})
}

// 添加
func (h *UserHandler) AddUserHandler(c *gin.Context)  {
	result := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := model.User{}
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": result})
		return
	}
	r, err := h.UserSrv.Add(u)
	if err != nil {
		result.Msg = err.Error()
		return
	}
	if r.UserId == "" {
		c.JSON(http.StatusOK, gin.H{"result": result})
		return
	}
	result.Code = int(enum.OperateOk)
	result.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK, gin.H{"result": result})
}