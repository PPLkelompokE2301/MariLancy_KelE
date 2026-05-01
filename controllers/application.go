package controllers

import (
	"net/http"

	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

// Author: Arga
// PBI: KF-05
// Sprint: Sprint 1
func ApplyJob(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		JobID uint `json:"job_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	var freelancer models.Freelancer
	if err := config.DB.First(&freelancer, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Freelancer tidak ditemukan"})
		return
	}
	if freelancer.Resume == "" || freelancer.Certificates == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lengkapi resume & certificates dulu"})
		return
	}

	var job models.Job
	if err := config.DB.First(&job, input.JobID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job tidak ditemukan"})
		return
	}
	if job.Status == "dihapus" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Job sudah dihapus"})
		return
	}
	if job.Status == "ditutup" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Job sudah ditutup"})
		return
	}

	var existing models.Application
	err := config.DB.Where("freelancer_id = ? AND job_id = ?", userID, input.JobID).First(&existing).Error

	if err == nil {

		if existing.Status == "withdrawn" {
			existing.Status = "pending"

			if err := config.DB.Save(&existing).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal re-apply"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Berhasil melamar ulang job"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Sudah pernah melamar job ini"})
		return
	}

	app := models.Application{
		FreelancerID: userID,
		JobID:        input.JobID,
		Status:       "pending",
	}

	if err := config.DB.Create(&app).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melamar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil melamar job"})
}
