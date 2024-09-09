package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	IP         string
	Port       string
	DataSource struct {
		Host     string
		Port     string
		Type     string
		Username string
		Password string
		Dbname   string
	}
}

var (
	conf *Config
	v    = viper.New()
)

func init() {
	v.AutomaticEnv()
	v.SetConfigType("yaml")
}
func GetConfig() *Config {
	if conf != nil {
		return conf
	}
	configFile := flag.String("c", "./config.yml", "Path to config file")
	return ReadConfig(*configFile)
}
func ReadConfig(path string) *Config {
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		GenerateConfig()
		panic("配置文件文件不存在: " + err.Error())
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic("配置文件反序列化失败: " + err.Error())
	}

	log.Println("配置文件内容加载成功: ", path)
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
	err := v.WriteConfig()
	if err != nil {
		return
	}
}
