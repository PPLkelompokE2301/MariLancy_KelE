// Author: Hanif
// PBI: KF-13
// Sprint: Sprint 1
package main

import (
	"marilancy/config"
	"marilancy/models"
	"marilancy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 2
	config.ConnectDB()

	// Author: Hanif
	// PBI: KF-13
	// Sprint: Sprint 1
	config.DB.AutoMigrate(
		&models.Admin{},
		&models.Client{},
		&models.Freelancer{},
	)

	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 1
	config.SeedAdmin()

	// Author: Hanif
	// PBI: KF-13
	// Sprint: Sprint 1
	r := gin.Default()
	r.Static("/static", "./static")
	r.Static("/uploads", "./uploads")
	r.LoadHTMLGlob("templates/*")

	routes.SetupRoutes(r)

	// Author: Hanif
	// PBI: KF-13
	// Sprint: Sprint 1
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "MariLancy base project is running",
		})
	})

	r.Run(":8080")
}
