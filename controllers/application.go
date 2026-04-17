// Author: Danu
// PBI: KF-06
// Sprint: Sprint 1
package controllers

import (
	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

func GetJobApplicants(c *gin.Context) {

	jobID := c.Param("id")

	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var job models.Job
	if err := config.DB.First(&job, jobID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Job tidak ditemukan"})
		return
	}

	if job.ClientID != userID {
		c.JSON(403, gin.H{"error": "Tidak bisa lihat pelamar job orang lain"})
		return
	}

	var apps []models.Application

	err := config.DB.
		Preload("Freelancer").
		Where("job_id = ?", jobID).
		Order("tanggal_daftar desc").
		Find(&apps).Error

	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil pelamar"})
		return
	}

	c.JSON(200, apps)
}
