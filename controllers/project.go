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

func UpdateTaskStatus(c *gin.Context) {
	taskID := c.Param("task_id")
	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&models.Task{}).Where("id = ?", taskID).Update("status", input.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update task"})
		return
	}

	var task models.Task
	if err := config.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Status task diperbarui"})
		return
	}

	var project models.Project
	if err := config.DB.Preload("Tasks").First(&project, task.ProjectID).Error; err == nil {

		totalTasks := len(project.Tasks)
		completedTasks := 0
		for _, t := range project.Tasks {
			if t.Status == "done" {
				completedTasks++
			}
		}

		newProgress := 0
		if totalTasks > 0 {
			newProgress = (completedTasks * 100) / totalTasks
		}

		newStatus := project.Status
		if newProgress > 0 && newProgress < 100 && project.Status == "active" {
			newStatus = "inprogress"
		} else if newProgress == 0 && project.Status == "inprogress" {
			newStatus = "active" 
		}

		config.DB.Model(&models.Project{}).Where("id = ?", project.ID).Updates(map[string]interface{}{
			"progress": newProgress,
			"status":   newStatus,
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status diperbarui"})
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
func UpdateTaskPriority(c *gin.Context) {
	taskID := c.Param("task_id")
	var input struct {
		Priority string `json:"priority"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Model(&models.Task{}).Where("id = ?", taskID).Update("priority", input.Priority).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal update prioritas"})
		return
	}

	c.JSON(200, gin.H{"message": "Prioritas diperbarui"})
}

