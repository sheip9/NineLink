package config

import (
	"github.com/sheip9/ninelink/internal/enum"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	IP         string
	Port       string
	DataSource struct {
		Host     string
		Port     string
		Type     enum.DBType
		Username string
		Password string
		Dbname   string
	}
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
	v.SetDefault("Port", "8080")
	v.SetDefault("DataSource.Host", "127.0.0.1")
	v.SetDefault("DataSource.Port", "3306")
	v.SetDefault("DataSource.Type", "mysql")
	v.SetDefault("DataSource.Username", "root")
	v.SetDefault("DataSource.Password", "123456")
	v.SetDefault("DataSource.Dbname", "ninelink")
	err := v.WriteConfig()
	if err != nil {
		return
	}
}
