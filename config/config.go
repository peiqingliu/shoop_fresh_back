package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const PageSize int = 10

type Config struct {
	Name string
}

func (c *Config) initConfig() error  {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	}else {
		viper.AddConfigPath("conf")  // 加载默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 设置文件格式
	err := viper.ReadInConfig()
	fmt.Println(viper.Get("database.name"))
	if err !=nil {
		logrus.Error("read config failed: %v", err)
		return err
	}
	return nil
}

func (c *Config) watchConfig()  {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Info("配置文件已经修改:",e.Name)
	})
}

func Init(name string) error {
	c := Config{
		Name:name,
	}
	if err := c.initConfig(); err != nil {
		return err
	}
	c.watchConfig()
	return nil
}