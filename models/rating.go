package models

import "time"

// Author: Aura
// PBI: KF-15
// Sprint: Sprint 2
type Rating struct {
	ID           uint `gorm:"primaryKey"`
	Nilai        float64
	Komentar     string
	ClientID     uint
	FreelancerID uint
	ProjectID    uint
	CreatedAt    time.Time
}
