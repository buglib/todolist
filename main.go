package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() (db *gorm.DB, err error) {
	dsn := "buglib:123456@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func main() {
	_, err := InitMysql()
	if err != nil {
		panic(err)
	}
}
