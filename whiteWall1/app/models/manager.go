package models

type Manager struct {
	ManagerID    uint   `json:"manager_id" gorm:"primary_key;AUTO_INCREMENT"`
	Account      string `json:"account"`
	Password     string `json:"-"`
	Name         string `json:"name"`
	ManagerState bool   `json:"manager_state"`
}
