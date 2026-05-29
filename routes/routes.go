	// Author: Aura
	// PBI: KF-11
	// Sprint: Sprint 1
	r.GET("/jobs/:id", controllers.GetJobDetail)

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

	// Author: Aura
	// PBI: KF-11
	// Sprint: Sprint 1
	r.GET("/job/detail", func(c *gin.Context) {
		c.HTML(200, "job_detail.html", nil)
	})

	// Author: Aura
	// PBI: KF-07
	// Sprint: Sprint 1
	r.GET("/client/saved-applicants", func(c *gin.Context) {
		c.HTML(200, "saved-applicants.html", nil)
	})

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

