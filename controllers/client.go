// Author: Fadhil
// PBI: KF-02
// Sprint: Sprint 1
package controllers

import (
	"strconv"

	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

func resolveClientID(c *gin.Context) (uint, error) {
	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	idParam := c.Query("id")
	if idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return 0, err
		}
		return uint(parsedID), nil
	}

	var client models.Client
	if err := config.DB.Order("id asc").First(&client).Error; err != nil {
		return 0, err
	}

	return client.ID, nil
}

func GetClientProfile(c *gin.Context) {
	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	userID, err := resolveClientID(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "Client tidak ditemukan"})
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
	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	userID, err := resolveClientID(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "Client tidak ditemukan"})
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
