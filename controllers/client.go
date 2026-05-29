// Author: Fadhil
// PBI: KF-02
// Sprint: Sprint 1
package controllers

import (
	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

func GetClientProfile(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var client models.Client
	if err := config.DB.First(&client, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Client tidak ditemukan"})
		return
	}

	c.JSON(200, client)
}

func UpdateClientProfile(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var client models.Client
	if err := config.DB.First(&client, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Client tidak ditemukan"})
		return
	}

	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	delete(input, "id")
	delete(input, "email")
	delete(input, "password")
	delete(input, "role")
	delete(input, "created_at")

	if err := config.DB.Model(&client).Updates(input).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal update profil"})
		return
	}

	c.JSON(200, gin.H{"message": "Profil berhasil diupdate"})
}

func GetApplicationsByClient(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var jobs []models.Job
	if err := config.DB.Where("client_id = ?", userID).Find(&jobs).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil job"})
		return
	}

	var apps []models.Application

	for _, job := range jobs {
		var jobApps []models.Application
		config.DB.Where("lowongan_id = ?", job.ID).Find(&jobApps)
		apps = append(apps, jobApps...)
	}

	c.JSON(200, apps)
}

func ApproveApplication(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")

	var app models.Application
	if err := config.DB.First(&app, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Application tidak ditemukan"})
		return
	}

	var job models.Job
	if err := config.DB.First(&job, app.JobID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Job tidak ditemukan"})
		return
	}

	if job.ClientID != userID {
		c.JSON(403, gin.H{"error": "Bukan job milikmu"})
		return
	}

	app.Status = "accepted"

	if err := config.DB.Save(&app).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal approve"})
		return
	}

	c.JSON(200, gin.H{"message": "Application approved"})
}

func RejectApplication(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")

	var app models.Application
	if err := config.DB.First(&app, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Application tidak ditemukan"})
		return
	}

	var job models.Job
	if err := config.DB.First(&job, app.JobID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Job tidak ditemukan"})
		return
	}

	if job.ClientID != userID {
		c.JSON(403, gin.H{"error": "Bukan job milikmu"})
		return
	}

	app.Status = "rejected"

	if err := config.DB.Save(&app).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal reject"})
		return
	}

	c.JSON(200, gin.H{"message": "Application rejected"})
}
