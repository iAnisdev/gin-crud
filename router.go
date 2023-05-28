package main

import "github.com/gin-gonic/gin"

func MainRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello from gins",
	})
}

var todos []Todo

func ListTodos(ctx *gin.Context) {

}

func AddTodo(ctx *gin.Context) {

}

func GetTodo(ctx *gin.Context) {

}

func DeleteTodos(ctx *gin.Context) {

}

func UpdateTodos(ctx *gin.Context) {

}
