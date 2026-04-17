// Author: Arga
// PBI: KF-17
// Sprint: Sprint 1
package models

import "time"

// Author: Arga
// PBI: KF-17
// Sprint: Sprint 1
type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	NamaAdmin string `gorm:"type:varchar(100)"`
	Email     string `gorm:"unique"`
	Password  string
	Role      string
	CreatedAt time.Time
}
