package repositories

import (
	"errors"

	"github.com/pwn922/auth-service/internal/models"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetByID(id string) (*models.Role, error) {
	var role models.Role
	if err := r.db.First(&role, "id = ?", id).Error; err != nil {
		return nil, errors.New("role not found")
	}
	return &role, nil
}

func (r *roleRepository) GetByName(roleName string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("role_name = ?", roleName).First(&role).Error; err != nil {
		return nil, errors.New("role not found")
	}
	return &role, nil
}

func (r *roleRepository) List() ([]*models.Role, error) {
	var roles []*models.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) Update(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) DeleteByID(id string) error {
	return r.db.Delete(&models.Role{}, id).Error
}