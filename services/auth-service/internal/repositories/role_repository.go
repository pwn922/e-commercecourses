package repositories

import (
	"github.com/pwn922/auth-service/internal/models"
)

type RoleRepository interface {
	Create(role *models.Role) error
	GetByID(id string) (*models.Role, error)
	GetByName(roleName string) (*models.Role, error)
	List() ([]*models.Role, error)
	Update(role *models.Role) error
	DeleteByID(id string) error
}