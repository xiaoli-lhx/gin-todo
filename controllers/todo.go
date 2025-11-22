package controllers

import (
	"gin-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /todo
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// POST /todo
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newTodo)
}

// PUT /todo/:id
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := models.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该ID"})
		return
	}
	todo.Status = !todo.Status
	models.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// DELETE /todo/:id
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
