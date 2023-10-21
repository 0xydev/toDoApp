package handler

import (
	"net/http"
	"toDoApp/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}
func GetTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&todo).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
func GetCompletedTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Where("completed = ?", true).Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func GetIncompleteTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Where("completed = ?", false).Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}
