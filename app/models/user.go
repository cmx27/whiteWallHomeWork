package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	Account      string `json:"account"`
	Password     string `json:"-"`
	Name         string `json:"name"`
	ManagerState bool   `json:"manager_state"`
}

type Student struct {
	StudentID    uint   `json:"student_id" gorm:"primary_key"`
	Sex          string `json:"sex"`
	Major        string `json:"major"`
	Account      string `json:"account"`
	Password     string `json:"-"`
	Name         string `json:"name"`
	ManagerState bool   `json:"manager_state"`
}