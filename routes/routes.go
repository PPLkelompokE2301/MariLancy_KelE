package routes

import (
	"marilancy/controllers"
	"marilancy/middleware"

	"github.com/gin-gonic/gin"
)

// Author: Hanif
// PBI: KF-13
// Sprint: Sprint 1
func SetupRoutes(r *gin.Engine) {

	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 1
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Author: Hanif
	// PBI: KF-17
	// Sprint: Sprint 2
	r.POST("/api/support/lapor", middleware.AuthMiddleware(), controllers.CreateSupportTicket)

	// Author: Fadhil
	// PBI: KF-04
	// Sprint: Sprint 1
	r.GET("/jobs", controllers.GetJobs)

	// Author: Aura
	// PBI: KF-11
	// Sprint: Sprint 1
	r.GET("/jobs/:id", controllers.GetJobDetail)

	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	r.GET("/clients/:id", controllers.GetClientByID)

	// Author: Aura
	// PBI: KF-15
	// Sprint: Sprint 2
	r.GET("/freelancer/:id/rating-summary", controllers.GetFreelancerRatingSummary)

	// Author: Aura
	// PBI: KF-15
	// Sprint: Sprint 2
	r.GET("/client/rating", func(c *gin.Context) {
		c.HTML(200, "rating.html", nil)
	})

	// Author: Danu
	// PBI: KF-10
	// Sprint: Sprint 2
	r.GET("/freelancer/projects", func(c *gin.Context) {
		c.HTML(200, "my-projects.html", nil)
	})

	// Author: Danu
	// PBI: KF-10
	// Sprint: Sprint 2
	r.GET("/client/projects", func(c *gin.Context) {
		c.HTML(200, "my-projects-client.html", nil)
	})

	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	r.GET("/lihatprofileclient.html", func(c *gin.Context) {
		c.HTML(200, "lihatprofileclient.html", nil)
	})

	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 2
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "dashboard_admin.html", nil)
	})

	// Author: Aura
	// PBI: KF-11
	// Sprint: Sprint 1
	r.GET("/job/detail", func(c *gin.Context) {
		c.HTML(200, "job_detail.html", nil)
	})

	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	r.GET("/client/profile/view", func(c *gin.Context) {
		c.HTML(200, "profile.html", nil)
	})

	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	r.GET("/freelancer/profile/view", func(c *gin.Context) {
		c.HTML(200, "profilefree.html", nil)
	})

	// Author: Danu
	// PBI: KF-06
	// Sprint: Sprint 1
	r.GET("/client/applicants", func(c *gin.Context) {
		c.HTML(200, "pendaftar.html", nil)
	})

	// Author: Aura
	// PBI: KF-07
	// Sprint: Sprint 1
	r.GET("/client/saved-applicants", func(c *gin.Context) {
		c.HTML(200, "saved-applicants.html", nil)
	})

	// Author: Danu
	// PBI: KF-16
	// Sprint: Sprint 2
	r.GET("/freelancer/notification", func(c *gin.Context) {
		c.HTML(200, "notification.html", nil)
	})

	// Author: Rania
	// PBI: KF-08
	// Sprint: Sprint 2
	r.GET("/freelancer/workspace", func(c *gin.Context) {
		c.HTML(200, "workspace-freelancer.html", nil)
	})

	// Author: Rania
	// PBI: KF-08
	// Sprint: Sprint 2
	r.GET("/client/workspace", func(c *gin.Context) {
		c.HTML(200, "workspace-client.html", nil)
	})

	// Author: Rania
	// PBI: KF-14
	// Sprint: Sprint 2
	r.GET("/chat", func(c *gin.Context) {
		c.HTML(200, "chat-list.html", nil)
	})

	// Author: Rania
	// PBI: KF-14
	// Sprint: Sprint 2
	r.GET("/chat/room", func(c *gin.Context) {
		c.HTML(200, "chat-room.html", nil)
	})

	// Author: Fadhil
	// PBI: KF-12
	// Sprint: Sprint 2
	r.GET("/client/transactions", func(c *gin.Context) {
		c.HTML(200, "transactions.html", nil)
	})

	// Author: Rania
	// PBI: KF-14
	// Sprint: Sprint 2
	apiChat := r.Group("/api/chat")
	apiChat.Use(middleware.AuthMiddleware())
	{
		apiChat.POST("/send", controllers.SendMessage)
		apiChat.GET("/history", controllers.GetChatHistory)
		apiChat.GET("/list", controllers.GetChatList)
		apiChat.POST("/read", controllers.MarkAsRead)
	}

	// Author: Rania
	// PBI: KF-08
	// Sprint: Sprint 2
	apiProjects := r.Group("/api/projects")
	apiProjects.Use(middleware.AuthMiddleware())
	{
		apiProjects.GET("/:id/rating", controllers.GetProjectRating)
		apiProjects.GET("/:id", controllers.GetProjectDetail)
		apiProjects.POST("/task", controllers.CreateTask)
		apiProjects.PUT("/task/:task_id", controllers.UpdateTaskStatus)
		apiProjects.PUT("/:id/complete", controllers.CompleteProject)
		apiProjects.PUT("/:id/revision", controllers.RequestRevision)
		apiProjects.PUT("/:id/cancel", controllers.CancelProject)
		apiProjects.DELETE("/task/:task_id", controllers.DeleteTask)
		apiProjects.PATCH("/task/:task_id/title", controllers.UpdateTaskTitle)
		apiProjects.PUT("/task/:task_id/priority", controllers.UpdateTaskPriority)
		apiProjects.PUT("/:id/deadline", controllers.UpdateProjectDeadline)
		apiProjects.POST("/:id/pay", controllers.ConfirmPayment)
		apiProjects.GET("/transactions", controllers.GetClientTransactions)
		apiProjects.PUT("/transactions/:id/status", controllers.UpdateTransactionStatus)
	}

	// Author: Fadhil
	// PBI: KF-05
	// Sprint: Sprint 1
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
		freelancer.GET("/my-projects", controllers.GetMyProjects)
	}

	// Author: Aura
	// PBI: KF-03
	// Sprint: Sprint 1
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
		client.GET("/my-projects", controllers.GetClientProjects)
		client.POST("/rating", controllers.CreateRating)
		client.GET("/check-rating", controllers.CheckRating)
	}

	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 2
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
		admin.PUT("/jobs/:id/restore", controllers.RestoreJobs)
		admin.GET("/transactions", controllers.AdminGetTransactions)
		admin.GET("/support-tickets", controllers.GetSupportTickets)

	}
}
