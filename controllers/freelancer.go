// Author: Fadhil
// PBI: KF-02
// Sprint: Sprint 1

// Author: Rania
// PBI: KF-09
// Sprint: Sprint 1
package controllers

import (
	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

func GetMyApplications(c *gin.Context) {

	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var apps []models.Application

	err := config.DB.
		Preload("Job").
		Where("freelancer_id = ?", userID).
		Order("tanggal_daftar desc").
		Find(&apps).Error

	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}

	for i := range apps {

		if apps[i].Job.ID == 0 && apps[i].JobID != 0 {

			var job models.Job
			err := config.DB.First(&job, apps[i].JobID).Error
			if err == nil {
				apps[i].Job = job
			}
		}

		if apps[i].JobID != 0 && apps[i].Job.ID == 0 {
			apps[i].Status = "ditolak (job tidak tersedia)"
			continue
		}

		if apps[i].Job.Status == "dihapus" {
			apps[i].Status = "ditolak (job dihapus)"
		}

		if apps[i].Job.Status == "ditutup" {
			apps[i].Status = "pending"
		}
	}

	c.JSON(200, apps)
}

func GetMyCompletedJobs(c *gin.Context) {
	id, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var apps []models.Application
	config.DB.Where("freelancer_id = ? AND status = ?", id, "accepted").Find(&apps)

	c.JSON(200, apps)
}
