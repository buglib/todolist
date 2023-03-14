package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	err := initMysql()
	if err != nil {
		panic(err)
	}

	router := initRouter()
	router.Run("0.0.0.0:8080")
}

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

func initMysql() (err error) {
	dsn := "buglib:123456@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Task{})
	return
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")
	router.Static("/static", "static")
	registerHandlers(router)
	return router
}

func registerHandlers(router *gin.Engine) {
	router.GET("/", getIndex)

	v1 := router.Group("v1")
	v1.POST("/todolist", postTodolist)
	v1.GET("/todolist", getTodolist)
	v1.PUT("/todolist/:id", putTodoItem)
	v1.DELETE("/todolist/:id", deleteTodoItem)
}

func getIndex(ctx *gin.Context) {

}

func postTodolist(ctx *gin.Context) {

}

func getTodolist(ctx *gin.Context) {

}

func putTodoItem(ctx *gin.Context) {

}

func deleteTodoItem(ctx *gin.Context) {

}
