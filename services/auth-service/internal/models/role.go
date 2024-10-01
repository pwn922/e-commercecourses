package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          string `gorm:"type:uuid;primary_key;" json:"id"`
	RoleName    string `gorm:"unique;not null" json:"role_name"`
	Description string `json:"description"`
	Users       []User `gorm:"foreignKey:RoleID" json:"users"`
}