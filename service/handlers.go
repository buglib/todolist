package service

import (
	"errors"
	"fmt"
	"todolist/domain"
	"todolist/infra/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetIndex(ctx *gin.Context) {

}

func PostTodolist(ctx *gin.Context) {
	var (
		task       domain.Task
		statusCode int
		respBody   interface{}
	)
	ctx.BindJSON(&task)
	err := db.Db.Create(&task).Error
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

func GetTodolist(ctx *gin.Context) {
	var (
		tasks      []domain.Task
		statusCode int
		respBody   interface{}
	)
	err := db.Db.Find(&tasks).Error
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

func PutTodoItem(ctx *gin.Context) {
	var (
		task       domain.Task
		statusCode int
		respBody   interface{}
		status     string
		msg        string
		data       interface{}
	)
	id, _ := ctx.Params.Get("id")
	err := db.Db.First(&task).Error
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
		temp := domain.Task{}
		ctx.BindJSON(&temp)
		// err = db.Db.Save(&task).Error
		err = db.Db.Model(&task).Updates(map[string]interface{}{"TaskInfo": temp.TaskInfo, "State": temp.State}).Error
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

func DeleteTodoItem(ctx *gin.Context) {
	var (
		task   domain.Task
		code   int
		status string
		msg    string
		data   interface{}
	)
	id, _ := ctx.Params.Get("id")
	// err := db.Db.Where("id = ?", id).Delete(Task{}).Error
	err := db.Db.First(&task, id).Error
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
		err = db.Db.Delete(&task, id).Error
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
