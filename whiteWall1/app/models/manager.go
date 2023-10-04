package models

type Manager struct {
	ManagerID uint `json:"manager_id" gorm:"primary_key"`
	User
}
