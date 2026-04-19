// Author: Fadhil
// PBI: KF-09
// Sprint: Sprint 1
package controllers

import (
	"fmt"
	"marilancy/config"
	"marilancy/models"

	"github.com/gin-gonic/gin"
)

func GetFreelancerProfile(c *gin.Context) {
	id, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	fmt.Println("🔥 HIT GET /freelancer/profile, ID:", id)

	var user models.Freelancer
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func UpdateFreelancerProfile(c *gin.Context) {
	id, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.Freelancer
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	user.Nama = c.PostForm("nama")
	user.Gender = c.PostForm("gender")
	user.Location = c.PostForm("location")
	user.EducationLevel = c.PostForm("education_level")
	user.JobInterest = c.PostForm("job_interest")
	user.Bio = c.PostForm("bio")
	user.Skill = c.PostForm("skill")
	user.WorkPre = c.PostForm("work_pre")

	fmt.Sscanf(c.PostForm("age"), "%d", &user.Age)
	fmt.Sscanf(c.PostForm("years_of_experience"), "%d", &user.YearsOfExperience)

	user.MonthlySalaryExp = c.PostForm("monthly_salary_exp")

	file, err := c.FormFile("resume")
	if err == nil {
		path := "uploads/resume_" + file.Filename
		c.SaveUploadedFile(file, path)
		user.Resume = "/" + path
	}
	fileCert, err := c.FormFile("certificates")
	if err == nil {
		path := "uploads/cert_" + fileCert.Filename
		c.SaveUploadedFile(fileCert, path)
		user.Certificates = "/" + path
	}

	fileAttach, err := c.FormFile("attachments")
	if err == nil {
		path := "uploads/attach_" + fileAttach.Filename
		c.SaveUploadedFile(fileAttach, path)
		user.Attachments = "/" + path
	}

	config.DB.Save(&user)

	c.JSON(200, gin.H{"msg": "Profile updated"})
}
