package utils

import (
	"fmt"
	"github.com/sheip9/ninelink/config"
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
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db = CreateDBInstance(ds.Type, ds.Username, ds.Password, ds.Host, ds.Port, ds.Dbname, gormConfig)
	return db
}

func CreateDBInstance(dbType string, username string, password string, host string, port string, dbName string, gormConfig *gorm.Config) *gorm.DB {
	var _db *gorm.DB
	switch dbType {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			username, password, host, port, dbName,
		)
		_db, _ = gorm.Open(mysql.Open(dsn), gormConfig)
	case "postgres":
		dsn := fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s",
			username, password, host, port, dbName,
		)
		_db, _ = gorm.Open(postgres.Open(dsn), gormConfig)
	default:
		_ = fmt.Errorf("")
		return nil
	}
	return _db
}
