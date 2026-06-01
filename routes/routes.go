	// Author: Fadhil
	// PBI: KF-12
	// Sprint: Sprint 2
	r.GET("/client/transactions", func(c *gin.Context) {
		c.HTML(200, "transactions.html", nil)
	})