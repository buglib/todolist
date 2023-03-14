package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义任务状态
const (
	Done = iota
	Todo
)

type Task struct {
	gorm.Model
	TaskInfo string
	State    uint
}

func InitMysql() (db *gorm.DB, err error) {
	dsn := "buglib:123456@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Task{})
	return
}

func main() {
	_, err := InitMysql()
	if err != nil {
		panic(err)
	}
}
