package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "buglib:123456@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	err := initMysql(dsn)
	if err != nil {
		panic(err)
	}

	router := initRouter()
	router.Run("0.0.0.0:8080")
}

// 定义任务状态
const (
	Todo = iota
	Done
)

type Task struct {
	ID       uint   `json:"id"`
	TaskInfo string `json:"taskInfo"`
	State    uint   `json:"state"`
}

func initMysql(dsn string) (err error) {
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
	var (
		task       Task
		statusCode int
		respBody   interface{}
	)
	ctx.BindJSON(&task)
	err := db.Create(&task).Error
	if err != nil {
		statusCode = 500
		respBody = gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("%v", err),
			"data":    nil,
		}
	} else {
		statusCode = 200
		respBody = gin.H{
			"status":  "done",
			"message": "Success to create task",
			"data":    task,
		}
	}
	ctx.JSON(statusCode, respBody)
}

func getTodolist(ctx *gin.Context) {
	var (
		tasks      []Task
		statusCode int
		respBody   interface{}
	)
	err := db.Find(&tasks).Error
	if err != nil {
		statusCode = 500
		respBody = gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("%v", err),
			"data":    nil,
		}
	} else {
		statusCode = 200
		respBody = gin.H{
			"status":  "done",
			"message": "Success to list all tasks",
			"data":    tasks,
		}
	}
	ctx.JSON(statusCode, respBody)
}

func putTodoItem(ctx *gin.Context) {
	var (
		task       Task
		statusCode int
		respBody   interface{}
		status     string
		msg        string
		data       interface{}
	)
	id, _ := ctx.Params.Get("id")
	err := db.First(&task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = 404
			status = "failed"
			msg = fmt.Sprintf("Task '%s' not found", id)
			data = nil
		} else {
			statusCode = 500
			status = "failed"
			msg = fmt.Sprintf("%v", err)
			data = nil
		}
	} else {
		ctx.BindJSON(&task)
		err = db.Save(&task).Error
		if err != nil {
			statusCode = 500
			status = "failed"
			msg = fmt.Sprintf("%v", err)
			data = nil
		} else {
			statusCode = 200
			status = "done"
			msg = fmt.Sprintf("Success to update task '%s'", id)
			data = task
		}
	}
	respBody = gin.H{
		"status":  status,
		"message": msg,
		"data":    data,
	}
	ctx.JSON(statusCode, respBody)
}

func deleteTodoItem(ctx *gin.Context) {
	var (
		task   Task
		code   int
		status string
		msg    string
		data   interface{}
	)
	id, _ := ctx.Params.Get("id")
	// err := db.Where("id = ?", id).Delete(Task{}).Error
	err := db.First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = 404
			status = "failed"
			msg = fmt.Sprintf("Task '%s' not found", id)
			data = nil
		} else {
			code = 500
			status = "failed"
			msg = fmt.Sprintf("%v", err)
			data = nil
		}
	} else {
		err = db.Delete(&task, id).Error
		if err != nil {
			code = 500
			status = "failed"
			msg = fmt.Sprintf("%v", err)
			data = nil
		} else {
			code = 200
			status = "done"
			msg = fmt.Sprintf("Success to delete task '%s'", id)
			data = nil
		}
	}
	ctx.JSON(
		code,
		gin.H{
			"status":  status,
			"message": msg,
			"data":    data,
		},
	)
}
