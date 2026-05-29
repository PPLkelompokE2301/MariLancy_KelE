package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)


func CreateTask(c *gin.Context) {
	var input struct {
		ProjectID uint   `json:"project_id"`
		Title     string `json:"title"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		ProjectID: input.ProjectID,
		Title:     input.Title,
		Status:    "todo",
	}

	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task ditambahkan", "task": task})
}


func DeleteTask(c *gin.Context) {
	taskID := c.Param("task_id")

	if err := config.DB.Delete(&models.Task{}, taskID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task berhasil dihapus"})
}

func UpdateTaskTitle(c *gin.Context) {
	taskID := c.Param("task_id")
	var input struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	if err := config.DB.Model(&models.Task{}).Where("id = ?", taskID).Update("title", input.Title).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate judul task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Judul task diperbarui"})
}


