package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello from gins",
	})
}

var todos []Todo

func ListTodos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"todos": todos,
	})
}

func AddTodo(ctx *gin.Context) {
	var newTodo Todo

	if err := ctx.BindJSON(&newTodo); err != nil {
		fmt.Print(err)
		return
	}

	newTodo.id = len(todos) + 1
	newTodo.isdone = false

	todos = append(todos, newTodo)
	ctx.IndentedJSON(http.StatusCreated, todos)
}

func GetTodo(ctx *gin.Context) {

}

func DeleteTodo(ctx *gin.Context) {

}

func UpdateTodo(ctx *gin.Context) {

}
