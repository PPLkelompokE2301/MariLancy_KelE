// Author: Hanif
// PBI: KF-13
// Sprint: Sprint 1
package routes

import (
	"marilancy/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Author: Hanif
	// PBI: KF-13
	// Sprint: Sprint 1
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "dashboard_admin.html", nil)
	})

	// Author: Hanif
	// PBI: KF-13
	// Sprint: Sprint 1
	r.GET("/client", func(c *gin.Context) {
		c.HTML(200, "dashboard_client.html", nil)
	})

	// Author: Hanif
	// PBI: KF-13
	// Sprint: Sprint 1
	r.GET("/freelancer", func(c *gin.Context) {
		c.HTML(200, "dashboard_freelancer.html", nil)
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

	// Author: Arga
	// PBI: KF-17
	// Sprint: Sprint 1
	r.GET("/admin/data", controllers.AdminDashboardData)
	r.GET("/admin/freelancers", controllers.GetFreelancers)
	r.GET("/admin/clients", controllers.GetClients)
	r.DELETE("/admin/freelancers/:id", controllers.DeleteFreelancer)
	r.DELETE("/admin/clients/:id", controllers.DeleteClient)

	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	r.GET("/client/profile", controllers.GetClientProfile)
	r.PUT("/client/profile", controllers.UpdateClientProfile)

	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	r.GET("/freelancer/profile", controllers.GetFreelancerProfile)
	r.PUT("/freelancer/profile", controllers.UpdateFreelancerProfile)
}
