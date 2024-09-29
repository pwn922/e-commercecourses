package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pwn922/auth-service/graph/model"
	models "github.com/pwn922/auth-service/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// ARREGLAR LO DE MODELS.ROLE
func (r *userRepository) Create(newUserInput *model.UserInput) (*models.User, error) {
	var role models.Role
	if err := r.db.Where("role_id = ?", newUserInput.RoleID).First(&role).Error; err != nil {
		return nil, fmt.Errorf("rol no encontrado: %v", err)
	}

	// Aquí se debería hacer el hash de la contraseña
	//	hashedPassword := hashPassword(newUserInput.Password)

	newUser := &models.User{
		ID:       uuid.New().String(),
		FirstName:    newUserInput.FirstName,
		MiddleName:   newUserInput.MiddleName,
		LastName:     newUserInput.LastName,
		Email:        newUserInput.Email,
		PasswordHash: newUserInput.Password, // HASHEAR ANTES LA CONTRASEÑA
		Role:       role,
	}

	if err := r.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
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

func (r *userRepository) Update(updateUserInput *model.UserInput, id int) (*models.User, error) {
	updateUser := &models.User{
		//I:           id,
		FirstName:    updateUserInput.FirstName,
		MiddleName:   updateUserInput.MiddleName,
		LastName:     updateUserInput.LastName,
		Email:        updateUserInput.Email,
		PasswordHash: updateUserInput.Password,
		//Role:       string(*newUserInput.Role),
		//		Phone:      *newUserInput.Phone,
		//Address: *newUserInput.Address,
	}

	err := r.db.Save(&updateUser).Error
	return updateUser, err
}

func (r *userRepository) DeleteByID(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
