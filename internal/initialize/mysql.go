package initialize

import (
	"fmt"
	"time"
	"tsf-cron/config"
	"tsf-cron/pkg/core/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetString("Mysql.Username"),
		config.GetString("Mysql.Passwd"),
		config.GetString("Mysql.Ip"),
		config.GetInt("Mysql.Port"),
		config.GetString("Mysql.Db"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Can not connect Mysql: %s", err)
	}

	if sqlDB, err := db.DB(); err != nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return db
}
