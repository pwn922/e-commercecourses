package repositories

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	models "github.com/pwn922/auth-service/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(newUser *models.User) (*models.User, error) {
	if err := r.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *userRepository) GetByID(id string) (*models.User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("invalid UUID format")
	}

	var user models.User
	if err := r.db.Preload("Role").First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) List() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByField(field string, value string) (*models.User, error) {
	var user models.User
	if err := r.db.Where(fmt.Sprintf("%s = ?", field), value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(updateUser *models.User) (*models.User, error) {
	err := r.db.Save(&updateUser).Error
	return updateUser, err
}

func (r *userRepository) DeleteByID(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}