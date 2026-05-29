package models

import "time"

type Project struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	JobID        uint       `gorm:"uniqueIndex:idx_freelancer_job_project" json:"job_id"`
	Job          Job        `gorm:"foreignKey:JobID" json:"job"`
	ClientID     uint       `json:"client_id"`
	Client       Client     `gorm:"foreignKey:ClientID" json:"client"`
	FreelancerID uint       `gorm:"uniqueIndex:idx_freelancer_job_project" json:"freelancer_id"`
	Freelancer   Freelancer `gorm:"foreignKey:FreelancerID" json:"freelancer"`
	Status       string     `gorm:"default:'active'" json:"status"`
	Progress     int        `gorm:"default:0" json:"progress"`
	Tasks        []Task     `gorm:"foreignKey:ProjectID" json:"tasks"`

	Transactions []Transaction `gorm:"foreignKey:ProjectID" json:"transactions"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`

	SubmissionLink string `json:"submission_link"`
	SubmissionFile string `json:"submission_file"`
	PaymentStatus  string `gorm:"default:'unpaid'" json:"payment_status"`
}

type Task struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProjectID uint      `json:"project_id"`
	Title     string    `json:"title"`
	Status    string    `gorm:"default:'todo'" json:"status"`
	Priority  string    `gorm:"default:'low'" json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
