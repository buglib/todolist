package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMysql(dsn string) (err error) {
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db.AutoMigrate(&Task{})
	return
}

func AutoMigrate(model interface{}) {
	Db.AutoMigrate(model)
}
