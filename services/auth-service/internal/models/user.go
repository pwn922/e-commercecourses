package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           string `gorm:"type:uuid;primary_key;" json:"id"`
	FirstName    string `gorm:"not null" json:"first_name"`
	MiddleName   string `gorm:"not null" json:"middle_name"`
	LastName     string `gorm:"not null" json:"last_name"`
	Email        string `gorm:"unique;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"password_hash"`
	RoleID       uint   `gorm:"not null" json:"role_id"`       // Agregar un campo para la relación
	Role         Role   `gorm:"foreignKey:RoleID" json:"role"` // Relación con Role
}
