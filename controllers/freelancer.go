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

func resolveFreelancerID(c *gin.Context) (uint, error) {
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

	var freelancer models.Freelancer
	if err := config.DB.Order("id asc").First(&freelancer).Error; err != nil {
		return 0, err
	}

	return freelancer.ID, nil
}

func GetFreelancerProfile(c *gin.Context) {
	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	id, err := resolveFreelancerID(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var user models.Freelancer
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func UpdateFreelancerProfile(c *gin.Context) {
	// Author: Fadhil
	// PBI: KF-02
	// Sprint: Sprint 1
	id, err := resolveFreelancerID(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
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
	user.MonthlySalaryExp = c.PostForm("monthly_salary_exp")

	if age, err := strconv.Atoi(c.PostForm("age")); err == nil {
		user.Age = age
	}
	if years, err := strconv.Atoi(c.PostForm("years_of_experience")); err == nil {
		user.YearsOfExperience = years
	}

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
