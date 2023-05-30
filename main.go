package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", MainRoute)

	router.GET("/todos", ListTodos)

	router.GET("/todo/:id", GetTodo)

	router.POST("/todo", AddTodo)

	router.PATCH("/todo/:id/complete", CompleteTodo)

	router.PATCH("/todo/:id", UpdateTodo)

	router.DELETE("/todo/:id", DeleteTodo)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": "PAGE_NOT_FOUND", "message": "404 page not found",
		})
	})

	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed",
		})
	})

	router.Run()
}
