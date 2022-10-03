package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)
import _ "gorm.io/gorm"

var (
	once     sync.Once
	instance *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		dsn := "root:@tcp(127.0.0.1:4000)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
		db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(time.Hour)
		instance = db
	})
	return instance
}
