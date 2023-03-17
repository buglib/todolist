package main

import (
	"todolist/domain"
	"todolist/infra/db"
	"todolist/ui"
)

func main() {
	dsn := "buglib:123456@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	err := db.InitMysql(dsn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Task{})
	router := ui.InitRouter()
	router.Run("0.0.0.0:8080")
}
