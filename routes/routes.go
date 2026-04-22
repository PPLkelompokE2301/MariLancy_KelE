// Author: Hanif
// PBI: KF-13
// Sprint: Sprint 1
package routes

import (
	"marilancy/controllers"
	"marilancy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/jobs", controllers.GetJobs)
	r.GET("/jobs/:id", controllers.GetJobDetail)

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "dashboard_admin.html", nil)
	})

	r.GET("/job/detail", func(c *gin.Context) {
		c.HTML(200, "job_detail.html", nil)
	})

	r.GET("/client/profile/view", func(c *gin.Context) {
		c.HTML(200, "profile.html", nil)
	})

	r.GET("/freelancer/profile/view", func(c *gin.Context) {
		c.HTML(200, "profilefree.html", nil)
	})
	r.GET("/client/applicants", func(c *gin.Context) {
		c.HTML(200, "pendaftar.html", nil)
	})
	r.GET("/freelancer/notification", func(c *gin.Context) {
		c.HTML(200, "notification.html", nil)
	})

	freelancer := r.Group("/freelancer")
	freelancer.Use(
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("freelancer"),
	)
	{
		freelancer.GET("/profile", controllers.GetFreelancerProfile)
		freelancer.PUT("/profile", controllers.UpdateFreelancerProfile)
		freelancer.POST("/apply", controllers.ApplyJob)
		freelancer.GET("/applications", controllers.GetMyApplications)
		freelancer.POST("/withdraw", controllers.WithdrawApplication)
	}

	client := r.Group("/client")
	client.Use(
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("client"),
	)
	{
		client.GET("/profile", controllers.GetClientProfile)
		client.PUT("/profile", controllers.UpdateClientProfile)
		client.POST("/jobs", controllers.CreateJob)
		client.PUT("/jobs/:id", controllers.UpdateJob)
		client.GET("/jobs/:id/applicants", controllers.GetJobApplicants)
		client.PUT("/application/status", controllers.UpdateApplicationStatus)
		client.DELETE("/jobs/:id", controllers.DeleteJob)
	}

	admin := r.Group("/admin")
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("admin"),
	)
	{
		admin.GET("/data", controllers.AdminDashboardData)

		admin.GET("/freelancers", controllers.GetFreelancers)
		admin.GET("/clients", controllers.GetClients)

		admin.DELETE("/freelancers/:id", controllers.DeleteFreelancer)
		admin.DELETE("/clients/:id", controllers.DeleteClient)

		admin.GET("/jobs", controllers.AdminGetJobs)
		admin.DELETE("/jobs/:id", controllers.DeleteJobs)
	}
}
