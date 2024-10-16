package utils

import (
	"fmt"
	"github.com/sheip9/ninelink/config"
	"github.com/sheip9/ninelink/internal/enum"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
	Db = &db
	c  = config.Conf
)

func InitDB() *gorm.DB {
	if db != nil {
		return db
	}
	ds := (*c).DataSource
	var gormConfig = &gorm.Config{
		Logger: logger.Default.LogMode((*config.Conf).GetGormMode()),
	}
	db = CreateDBInstance(ds.Type, ds.Username, ds.Password, ds.Host, ds.Port, ds.DbName, gormConfig)
	return db
}

func CreateDBInstance(dbType enum.DBType, username string, password string, host string, port uint16, dbName string, gormConfig *gorm.Config) *gorm.DB {
	var _db *gorm.DB
	switch dbType {
	case enum.MySQL:
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			username, password, host, port, dbName,
		)
		_db, _ = gorm.Open(mysql.Open(dsn), gormConfig)
	case enum.Postgres:
		dsn := fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s",
			username, password, host, port, dbName,
		)
		_db, _ = gorm.Open(postgres.Open(dsn), gormConfig)
	default:
		_ = fmt.Errorf("数据库类型错误，可用值：%s, %s", enum.MySQL, enum.Postgres)
		return nil
	}
	return _db
}
