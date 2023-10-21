package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"toDoApp/handler"
	"toDoApp/models"
)

var err error

func main() {
	models.DB, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer models.DB.Close()

	models.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	r.GET("/todos", handler.GetTodos)
	r.POST("/todos", handler.CreateTodo)
	r.DELETE("/todos/:id", handler.DeleteTodo)
	r.GET("/todos/:id", handler.GetTodo)
	r.PUT("/todos/:id", handler.UpdateTodo)
	r.Run()
}
