package initialization

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shoop_Fresh_back/config"
	"shoop_Fresh_back/handler"
	"shoop_Fresh_back/model"
	"shoop_Fresh_back/repository"
	"shoop_Fresh_back/service"
)

var (
	DB *gorm.DB
	UserHandler handler.UserHandler
)

// 1、初始化配置文件
func initViper()  {
	err := config.Init("")
	if err != nil {
		logrus.Error("初始化配置文件错误")
	}
}
// 2、初始化数据库
func initDB()  {
	logrus.Info("start:初始化数据库")
	var err error
	conf := &model.DBConf{
		Host: viper.GetString("database.host"),
		User: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}
	// 格式化字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		true,
		"Local")
	// 此处为赋值 而不是:=
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("链接数据库失败",dsn)
	}
	logrus.Info("end:初始化数据库结束")
}

func initHandler()  {
	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: DB,
			},
		}}
}

// 初始化方法
func init()  {
	fmt.Println("初始化")
	initViper()
	initDB()
	initHandler()
}
