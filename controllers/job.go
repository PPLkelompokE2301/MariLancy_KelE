// Author: Fadhil
// PBI: KF-03
// Sprint: Sprint 1
package controllers

import (
	"encoding/json"
	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

func GetJobs(c *gin.Context) {
	var jobs []models.Job

	if err := config.DB.
		Preload("Client").
		Find(&jobs).Error; err != nil {

		c.JSON(500, gin.H{"error": "Gagal mengambil jobs"})
		return
	}

	for i := range jobs {
		var count int64

		err := config.DB.Model(&models.Application{}).
			Where("job_id = ?", jobs[i].ID).
			Count(&count).Error

		if err != nil {
			count = 0
		}

		jobs[i].ApplicationsCount = count
	}

	var result []gin.H

	for _, job := range jobs {

		var tags []string
		if job.Tags != "" {
			_ = json.Unmarshal([]byte(job.Tags), &tags)
		}

		result = append(result, gin.H{
			"id":                 job.ID,
			"judul":              job.Judul,
			"job_desc":           job.JobDesc,
			"kebutuhan_proyek":   job.KebutuhanProyek,
			"kebutuhan_skill":    job.KebutuhanSkill,
			"status":             job.Status,
			"kategori":           job.Kategori,
			"budget":             job.Budget,
			"batas_pendidikan":   job.BatasPendidikan,
			"pengalaman_kerja":   job.PengalamanKerja,
			"tipe":               job.Tipe,
			"lokasi_pelaksanaan": job.LokasiPelaksanaan,
			"tags":               tags,

			"share_job": job.ShareJob,
			"level":     job.Level,

			"client_id": job.ClientID,

			"client": gin.H{
				"id":          job.Client.ID,
				"nama_client": job.Client.NamaClient,
			},

			"applications_count": job.ApplicationsCount,
			"created_at":         job.CreatedAt,
		})
	}

	c.JSON(200, result)
}
