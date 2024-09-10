package models

import "gorm.io/gorm"

type User struct {
	gorm.Model // Incluye campos como ID, CreatedAt, UpdatedAt, DeletedAt
	ID         int    `gorm:"primary key;autoIncrement" json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}