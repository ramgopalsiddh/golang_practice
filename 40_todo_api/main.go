package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO structure
type todo struct {
	ID 			string `json:"id"`
	Item 		string `json:"item"`
	Completed 	bool `json:"completed"`
}

// Add fix data of todo
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

// function for Get all todos
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

// Function for Add a new Todo
func addTodo(context *gin.Context){
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// Get todo's data
func getTodo(context *gin.Context){
	// extract id
	id := context.Param("id")
	// extract todo data by use id
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todos not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

// Function for update/toggle todo's Status
func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todos not found"})
		return
	}

	// toggle completed option
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

// Function for get todo by use id
func getTodoById(id string) (*todo, error){
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}


func main(){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todo/:id", getTodo)
	router.PATCH("/todo/:id", toggleTodoStatus)
	router.POST("/todo",addTodo)
	router.Run("localhost:9090")
}