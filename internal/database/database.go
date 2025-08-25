package database

import (
	"fmt"
	"log"

	"github.com/sheip9/ninelink/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dB *gorm.DB
)

func GetDB() *gorm.DB {
	if dB == nil {
		panic("database not initialized")
	}
	return dB
}

// InitDB 初始化数据库连接
func InitDB() {
	if dB != nil {
		return
	}

	// 获取配置
	conf := config.Conf
	ds := conf.DataSource

	var gormConfig = &gorm.Config{
		Logger: logger.Default.LogMode(conf.GetGormMode()),
	}

	var err error
	dB, err = CreateDBInstance(ds.Type, ds.Username, ds.Password, ds.Host, ds.Port, ds.DbName, gormConfig)

	if err != nil {
		log.Fatalln(err)
		return
	}

	if err := dB.AutoMigrate(&Record{}); err != nil {
		log.Fatalln(err)
	}

}

// CreateDBInstance 创建数据库实例
func CreateDBInstance(dbType config.DBType, username string, password string, host string, port uint16, dbName string, gormConfig *gorm.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch dbType {
	case config.MySQL:
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, dbName,
		)
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)

	case config.Postgres:
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			host, username, password, dbName, port,
		)
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	default:
		err = fmt.Errorf("unsupported database type: %v", dbType)
	}

	return db, err
}
