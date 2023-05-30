package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MainRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello from gins",
	})
}

type body struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos []Todo = []Todo{}

func ListTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"todos": []interface{}{todos}})
}

func AddTodo(ctx *gin.Context) {
	body := body{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newTodo := Todo{
		Id:          len(todos) + 1,
		Title:       body.Title,
		Description: body.Description,
		Isdone:      false,
	}
	todos = append(todos, newTodo)
	ctx.JSON(http.StatusCreated, newTodo)
}

func GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	targetId, _ := strconv.ParseInt(id, 2, 0)
	targetTodo, index := Todo{}, -1
	for i := range todos {
		if todos[i].Id == int(targetId) {
			targetTodo = todos[i]
			index = i
		}
	}

	if index == -1 {
		message := fmt.Sprintf("No todo found with id %v ", id)
		ctx.JSON(http.StatusNotFound, gin.H{"message": message})
		return
	}
	ctx.JSON(http.StatusFound, gin.H{"todo": targetTodo})
}

func UpdateTodo(ctx *gin.Context) {

}

func CompleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	targetId, _ := strconv.ParseInt(id, 2, 0)
	targetTodo, index := Todo{}, -1
	for i := range todos {
		if todos[i].Id == int(targetId) {
			todos[i].Isdone = true
			index = i
			targetTodo = todos[i]
		}
	}
	if index == -1 {
		message := fmt.Sprintf("No todo found with id %v ", id)
		ctx.JSON(http.StatusNotFound, gin.H{"message": message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"todo": targetTodo})
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	targetId, _ := strconv.ParseInt(id, 2, 0)
	updatedTodos, index := []Todo{}, -1
	for i := range todos {
		if todos[i].Id == int(targetId) {
			index = i
		} else {
			updatedTodos = append(updatedTodos, todos[i])
		}
	}
	todos = updatedTodos
	if index == -1 {
		message := fmt.Sprintf("No todo found with id %v ", id)
		ctx.JSON(http.StatusNotFound, gin.H{"message": message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusGone, "message": "Todo deleted successfully"})
}
