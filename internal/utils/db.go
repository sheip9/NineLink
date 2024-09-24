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
	c  = config.Conf
)

func GetDB() *gorm.DB {
	ds := (*c).DataSource
	if db != nil {
		return db
	}
	var gormConfig = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	var _db *gorm.DB
	switch (*c).DataSource.Type {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			ds.Username, ds.Password, ds.Host, ds.Port, ds.Dbname,
		)
		_db, _ = gorm.Open(mysql.Open(dsn), gormConfig)
	case "postgres":
		dsn := fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s TimeZone=Asia/Shanghai",
			ds.Username, ds.Password, ds.Host, ds.Port, ds.Dbname,
		)
		_db, _ = gorm.Open(postgres.Open(dsn), gormConfig)
	default:
		panic("")
	}
	//_ = _db.AutoMigrate(&entity.Record{})
	db = _db
	return _db
}
