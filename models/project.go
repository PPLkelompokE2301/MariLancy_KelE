// Author: Fadhil
// PBI: KF-12
// Sprint: Sprint 2
type Transaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	ProjectID     uint      `json:"project_id"`
	Project       Project   `gorm:"foreignKey:ProjectID" json:"project"`
	ClientID      uint      `json:"client_id"`
	FreelancerID  uint      `json:"freelancer_id"`
	Nominal       float64   `json:"nominal"`
	BuktiTransfer string    `gorm:"type:varchar(255)" json:"bukti_transfer"`
	Status        string    `gorm:"default:'pending'" json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
