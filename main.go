package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"shoop_Fresh_back/initialization"
	"shoop_Fresh_back/middle"
)

func main()  {

	r := gin.Default()
	r.Use(middle.Cros())
	gin.SetMode(viper.GetString("mode"))

	user := r.Group("/api/user")
	{
		user.GET("/list",initialization.UserHandler.UserListHandler)
		user.POST("/add",initialization.UserHandler.AddUserHandler)
	}

	port := viper.GetString("port")
	err := r.Run(port)
	if err != nil {
		logrus.Error("启动失败")
	}
	
}

