package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/internal/enum"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
	"log"
)

type Config struct {
	IP         string
	Port       string
	Debug      bool
	DataSource struct {
		Host     string
		Port     uint16
		Type     enum.DBType
		Username string
		Password string
		DbName   string
	}
}

func (c Config) GetGinMode() string {
	if c.Debug {
		return gin.DebugMode
	}
	return gin.ReleaseMode
}
func (c Config) GetGormMode() logger.LogLevel {
	if c.Debug {
		return logger.Info
	}
	return logger.Error
}

var (
	conf *Config
	Conf = &conf
	v    = viper.New()
	File string
)

func init() {
	v.AutomaticEnv()
	v.SetConfigType("yaml")
}

func ReadConfig() *Config {
	v.SetConfigFile(File)
	if err := v.ReadInConfig(); err != nil {
		GenerateConfig()
		panic("配置文件文件不存在: " + err.Error())
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic("配置文件反序列化失败: " + err.Error())
	}

	log.Println("配置文件内容加载成功: ", File)
	return conf
}
func GenerateConfig() {
	v.SetDefault("IP", "")
	v.SetDefault("Port", 8080)
	v.SetDefault("Debug", false)
	v.SetDefault("DataSource.Host", "127.0.0.1")
	v.SetDefault("DataSource.Port", 3306)
	v.SetDefault("DataSource.Type", "mysql")
	v.SetDefault("DataSource.Username", "root")
	v.SetDefault("DataSource.Password", "123456")
	v.SetDefault("DataSource.DbName", "ninelink")
	err := v.WriteConfig()
	if err != nil {
		return
	}
}
