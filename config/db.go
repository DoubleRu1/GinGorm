package config

import (
	"GinGormCRUD/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func initDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := Appconfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initilize database, got an error %v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(Appconfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(Appconfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		log.Fatalf("Failed to initilize database, got an error %v", err)
	}
	global.DB = db
}
