
// Author: Arga
// PBI: KF-17
// Sprint: Sprint 1
package controllers

import (
	"net/http"

	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

func AdminDashboardData(c *gin.Context) {
	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 2
	var totalAdmins int64
	var totalFreelancers int64
	var totalClients int64

	config.DB.Model(&models.Admin{}).Count(&totalAdmins)
	config.DB.Model(&models.Freelancer{}).Count(&totalFreelancers)
	config.DB.Model(&models.Client{}).Count(&totalClients)

	c.JSON(http.StatusOK, gin.H{
		"admins":      totalAdmins,
		"freelancers": totalFreelancers,
		"clients":     totalClients,
	})
}

func GetFreelancers(c *gin.Context) {
	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 1
	var users []models.Freelancer

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil freelancer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetClients(c *gin.Context) {
	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 1
	var clients []models.Client

	if err := config.DB.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil client"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": clients,
	})
}

func DeleteFreelancer(c *gin.Context) {
	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 1
	id := c.Param("id")

	var user models.Freelancer
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Freelancer tidak ditemukan"})
		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Freelancer berhasil dihapus",
	})
}

func DeleteClient(c *gin.Context) {
	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 1
	id := c.Param("id")

	var client models.Client
	if err := config.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client tidak ditemukan"})
		return
	}

	config.DB.Delete(&client)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Client berhasil dihapus",
	})
}
