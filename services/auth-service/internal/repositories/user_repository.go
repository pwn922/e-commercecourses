package repositories

import (
	models "github.com/pwn922/auth-service/internal/models"
)

type UserRepository interface {
	Create(newUser *models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
	List() ([]*models.User, error)
	//GetByEmail(email string) (*models.User, error)
	GetByField(field string, value string) (*models.User, error)
	Update(updateUser *models.User) (*models.User, error)
	DeleteByID(id string) error
}
