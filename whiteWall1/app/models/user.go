package models

type User struct {
	UserID       uint   `json:"user_id" gorm:"AUTO_INCREMENT"`
	Account      string `json:"account"`
	Password     string `json:"-"`
	Name         string `json:"name"`
	ManagerState bool   `json:"manager_state"`
}

type Student struct {
	StudentID uint `json:"student_id" gorm:"primary_key"`
	User
	Sex   string `json:"sex"`
	Major string `json:"major"`
}
