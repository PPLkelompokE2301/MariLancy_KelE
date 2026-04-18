// Author: Aura
// PBI: KF-03
// Sprint: Sprint 1
package controllers

import (
	"encoding/json"
	"marilancy/config"
	"marilancy/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserID(c *gin.Context) (uint, bool) {
	val, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	switch v := val.(type) {
	case float64:
		return uint(v), true
	case int:
		return uint(v), true
	case uint:
		return v, true
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, false
		}
		return uint(i), true
	default:
		return 0, false
	}
}

func CreateJob(c *gin.Context) {
	var job models.Job

	var input struct {
		Judul             string   `json:"judul"`
		JobDesc           string   `json:"job_desc"`
		KebutuhanProyek   string   `json:"kebutuhan_proyek"`
		KebutuhanSkill    string   `json:"kebutuhan_skill"`
		Status            string   `json:"status"`
		Kategori          string   `json:"kategori"`
		Budget            string   `json:"budget"`
		BatasPendidikan   string   `json:"batas_pendidikan"`
		PengalamanKerja   string   `json:"pengalaman_kerja"`
		Tipe              string   `json:"tipe"`
		LokasiPelaksanaan string   `json:"lokasi_pelaksanaan"`
		Tags              []string `json:"tags"`
		ShareJob          string   `json:"share_job"`
		Level             string   `json:"level"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	job.Judul = input.Judul
	job.JobDesc = input.JobDesc
	job.KebutuhanProyek = input.KebutuhanProyek
	job.KebutuhanSkill = input.KebutuhanSkill
	job.Status = input.Status
	job.Kategori = input.Kategori
	job.Budget = input.Budget
	job.BatasPendidikan = input.BatasPendidikan
	job.PengalamanKerja = input.PengalamanKerja
	job.Tipe = input.Tipe
	job.LokasiPelaksanaan = input.LokasiPelaksanaan
	job.ShareJob = input.ShareJob
	job.Level = input.Level
	job.ClientID = userID

	tagsJSON, _ := json.Marshal(input.Tags)
	job.Tags = string(tagsJSON)

	if err := config.DB.Create(&job).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Job created"})
}

// Author: Aura
// PBI: KF-11
// Sprint: Sprint 1

func GetJobDetail(c *gin.Context) {
	id := c.Param("id")

	var job models.Job
	if err := config.DB.Preload("Client").First(&job, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Job tidak ditemukan"})
		return
	}

	c.JSON(200, job)
}

// Author: Danu
// PBI: KF-03 - Mengelola Proyek
// Sprint: Sprint 1

func DeleteJob(c *gin.Context) {
	id := c.Param("id")

	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var job models.Job
	if err := config.DB.First(&job, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Job tidak ditemukan"})
		return
	}

	if job.ClientID != userID {
		c.JSON(403, gin.H{"error": "Bukan job milikmu"})
		return
	}

	job.Status = "dihapus"

	if err := config.DB.Save(&job).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal menghapus job"})
		return
	}

	c.JSON(200, gin.H{"message": "Job dihapus (soft delete)"})
}

func UpdateJob(c *gin.Context) {
	id := c.Param("id")

	userID, ok := getUserID(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var job models.Job
	if err := config.DB.First(&job, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Job tidak ditemukan"})
		return
	}

	if job.ClientID != userID {
		c.JSON(403, gin.H{"error": "Bukan job milikmu"})
		return
	}

	var input struct {
		Judul             string `json:"judul"`
		JobDesc           string `json:"job_desc"`
		KebutuhanProyek   string `json:"kebutuhan_proyek"`
		KebutuhanSkill    string `json:"kebutuhan_skill"`
		Budget            string `json:"budget"`
		LokasiPelaksanaan string `json:"lokasi_pelaksanaan"`
		Status            string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if job.Status == "dihapus" {
		c.JSON(400, gin.H{"error": "Job sudah dihapus"})
		return
	}

	updateData := map[string]interface{}{
		"judul":              input.Judul,
		"job_desc":           input.JobDesc,
		"kebutuhan_proyek":   input.KebutuhanProyek,
		"kebutuhan_skill":    input.KebutuhanSkill,
		"budget":             input.Budget,
		"lokasi_pelaksanaan": input.LokasiPelaksanaan,
		"status":             input.Status,
	}

	if err := config.DB.Model(&models.Job{}).
		Where("id = ?", id).
		Updates(updateData).Error; err != nil {

		c.JSON(500, gin.H{"error": "Gagal update job"})
		return
	}

	c.JSON(200, gin.H{"message": "Job updated"})
}





