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

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.Freelancer{},
		&models.Client{},
		&models.Admin{},
		&models.Job{},
		&models.Application{},
		&models.Rating{},
	)
	config.SeedAdmin()

	r := gin.Default()

	r.Static("/static", "./static")
	r.Static("/uploads", "./uploads")
	r.LoadHTMLGlob("templates/*")

	routes.SetupRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	r.GET("/freelancer", func(c *gin.Context) {
		c.HTML(200, "dashboard_freelancer.html", nil)
	})

	r.GET("/client", func(c *gin.Context) {
		c.HTML(200, "dashboard_client.html", nil)
	})

	r.Run(":8080")
}
