package ui

import (
	"time"
	"todolist/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//AllowOrigins:    []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:    []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type"},
		AllowAllOrigins: true,
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.LoadHTMLFiles("templates/index.html")
	router.Static("/static", "static")
	registerHandlers(router)
	return router
}

func registerHandlers(router *gin.Engine) {
	router.GET("/", service.GetIndex)

	v1 := router.Group("v1")
	v1.POST("/todolist", service.PostTodolist)
	v1.GET("/todolist", service.GetTodolist)
	v1.PUT("/todolist/:id", service.PutTodoItem)
	v1.DELETE("/todolist/:id", service.DeleteTodoItem)
}
