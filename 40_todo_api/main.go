package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// TODO structure
type todo struct {
	ID        int    `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var db *sql.DB

func initDB() error {
	var err error
	// Open database connection
	db, err = sql.Open("postgres", "postgres://ram@localhost/todo_golang?sslmode=disable")
	if err != nil {
		return err
	}
	// Ping to check if database connection is successful
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}


// function for Get all todos
func getTodos(c *gin.Context) {
	// Retrieve all todos from database
	rows, err := db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var todos []todo
	// Loop through the rows and scan each into a todo struct
	for rows.Next() {
		var t todo
		err := rows.Scan(&t.ID, &t.Item, &t.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, t)
	}
	// Return todos as JSON
	c.JSON(http.StatusOK, todos)
}


// Function for Add a new Todo
func addTodo(c *gin.Context) {
	var newTodo todo

	// Bind JSON data to newTodo struct
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert new todo into database and retrieve its auto-generated ID
	err := db.QueryRow("INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id", newTodo.Item, newTodo.Completed).Scan(&newTodo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the newly created todo as JSON
	c.IndentedJSON(http.StatusCreated, newTodo)
}


// function for retrieve todo's data
func getTodo(c *gin.Context) {
	// Get todo ID from URL parameter
	id := c.Param("id")
	// Retrieve todo by its ID
	todo, err := getTodoByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Return the todo as JSON
	c.IndentedJSON(http.StatusOK, todo)
}


// Function for update/toggle todo's Status
func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Toggle completed status
	todo.Completed = !todo.Completed

	// Update todo in the database
	_, err = db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", todo.Completed, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}


// Function for get todo by use id
func getTodoByID(id string) (*todo, error) {
	var t todo
	// Query database for todo by its ID
	err := db.QueryRow("SELECT id, title, completed FROM todos WHERE id = $1", id).Scan(&t.ID, &t.Item, &t.Completed)
	if err != nil {
		// If todo not found, return an error
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return &t, nil
}


// function for delete todo's by id
func deleteTodoByID(c *gin.Context) {
	// Get todo ID from URL parameter
	id := c.Param("id")
	// Execute delete query
	result, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if any rows were affected by the delete operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If no rows were affected, return "Todo not found" message
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}


// main function
func main() {
	// Initialize database connection
	err := initDB()
	if err != nil {
		panic(err)
	}

	// Create a new router
	router := gin.Default()
	// Define API endpoints
	router.GET("/todos", getTodos)
	router.GET("/todo/:id", getTodo)
	router.PATCH("/todo/:id", toggleTodoStatus)
	router.POST("/todo", addTodo)
	router.DELETE("/todo/:id", deleteTodoByID)
	// Run the server
	if err := router.Run("localhost:9090"); err != nil {
		panic(err)
	}
}
