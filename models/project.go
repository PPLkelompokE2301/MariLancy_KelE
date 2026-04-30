package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

// Author: Rania
// PBI: KF-08
// Sprint: Sprint 1
func GetProjectDetail(c *gin.Context) {
	projectID := c.Param("id")
	var project models.Project

	if err := config.DB.Preload("Job").Preload("Client").Preload("Freelancer").Preload("Tasks").Where("id = ?", projectID).First(&project).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project tidak ditemukan"})
		return
	}

	totalTasks := len(project.Tasks)
	completedTasks := 0
	for _, task := range project.Tasks {
		if task.Status == "done" {
			completedTasks++
		}
	}

	progress := 0
	if totalTasks > 0 {
		progress = (completedTasks * 100) / totalTasks
	}

	c.JSON(http.StatusOK, gin.H{
		"project":  project,
		"progress": progress,
	})
}