package controllers

import (
	"marilancy/config"
	"marilancy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Author: Arga
// PBI: KF-17
// Sprint: Sprint 2
func AdminDashboardData(c *gin.Context) {
	var totalFreelancers, totalClients, totalJobs int64
	config.DB.Model(&models.Freelancer{}).Count(&totalFreelancers)
	config.DB.Model(&models.Client{}).Count(&totalClients)
	config.DB.Model(&models.Job{}).Where("status != ?", "dihapus").Count(&totalJobs)

	c.JSON(http.StatusOK, gin.H{
		"freelancers": totalFreelancers,
		"clients":     totalClients,
		"jobs":        totalJobs,
	})
}

// Author: Arga
// PBI: KF-17
// Sprint: Sprint 2
func AdminGetTransactions(c *gin.Context) {
	var results []struct {
		ProjectID      uint    `json:"project_id"`
		FreelancerName string  `json:"freelancer_name"`
		ClientName     string  `json:"client_name"`
		Nominal        float64 `json:"nominal"`
		Status         string  `json:"status"`
	}

	err := config.DB.Table("transactions").
		Select("transactions.project_id, freelancers.nama as freelancer_name, clients.nama_client as client_name, transactions.nominal, transactions.status").
		Joins("join jobs on jobs.id = transactions.project_id").
		Joins("join freelancers on freelancers.id = transactions.freelancer_id").
		Joins("join clients on clients.id = transactions.client_id").
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data transaksi: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
