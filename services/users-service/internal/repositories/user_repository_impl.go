package repositories

import (
	"github.com/pwn922/users-service/graph/model"
	models "github.com/pwn922/users-service/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(userInput *model.UserInput) (*models.User, error) {
	newUser := &models.User{
        Name:     *userInput.Name,
        Email:    *userInput.Email,
        Password: *userInput.Password,
    }

	err := r.db.Create(&newUser).Error
    return newUser, err
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
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

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(userInput *model.UserInput, id int) (*models.User, error) {
	updateUser := &models.User{
        ID:        id,
        Name:     *userInput.Name,
        Email:    *userInput.Email,
        Password: *userInput.Password,
    }

	err := r.db.Save(&updateUser).Error
	return updateUser, err
}

func (r *userRepository) DeleteByID(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}