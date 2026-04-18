// Author: Aura
// PBI: KF-07

// Author: Danu
// PBI: KF-06
// Sprint: Sprint 1
package controllers

import (
	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)
// Author: Aura
// PBI: KF-07
// Sprint: Sprint 1
func UpdateApplicationStatus(c *gin.Context) {

	var input struct {
		ApplicationID uint   `json:"application_id"`
		Status        string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	validStatus := map[string]bool{
		"pending":  true,
		"accepted": true,
		"rejected": true,
	}

	if !validStatus[input.Status] {
		c.JSON(400, gin.H{"error": "Status tidak valid"})
		return
	}

	var app models.Application

	if err := config.DB.First(&app, input.ApplicationID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Application not found"})
		return
	}

	if app.Status == "accepted" || app.Status == "rejected" {
		c.JSON(400, gin.H{"error": "Status sudah final, tidak bisa diubah"})
		return
	}

	app.Status = input.Status

	if err := config.DB.Save(&app).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal update status"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Status updated",
		"data":    app,
	})
}
