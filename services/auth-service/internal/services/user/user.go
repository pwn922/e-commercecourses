package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/pwn922/auth-service/graph/model"
	"github.com/pwn922/auth-service/internal/models"
	"github.com/pwn922/auth-service/internal/repositories"
	"github.com/pwn922/auth-service/internal/utils"
)

type UserService struct {
	UserRepository repositories.UserRepository
	RoleRepository repositories.RoleRepository
}

func NewUserService(userRepo repositories.UserRepository, roleRepo repositories.RoleRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
		RoleRepository: roleRepo,
	}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]*models.User, error) {
	panic("Not Implemented")
	//return s.UserRepository.List()
}

func (s *UserService) CreateUser(ctx context.Context, registerUserInput *models.RegisterUserInput) (*models.User, error) {
	role, err := s.RoleRepository.GetByName(registerUserInput.Role)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(registerUserInput.Password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		ID:             uuid.New().String(),
		FirstName:      registerUserInput.FirstName,
		MiddleName:     registerUserInput.MiddleName,
		LastName:       registerUserInput.LastName,
		Email:          registerUserInput.Email,
		HashedPassword: hashedPassword,
		RoleID:        role.ID,
	}

	createdUser, err := s.UserRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}


func (s *UserService) UpdateUser(ctx context.Context, updateUserInput *model.UserInput, id string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(updateUserInput.Password)
	if err != nil {
		return nil, err
	}

	updateUser := &models.User{
		ID:           id,
		FirstName:    updateUserInput.FirstName,
		MiddleName:   updateUserInput.MiddleName,
		LastName:     updateUserInput.LastName,
		Email:        updateUserInput.Email,
		HashedPassword: hashedPassword,
	}

	updatedUser, err := s.UserRepository.Update(updateUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}



func (s *UserService) DeleteUser(ctx context.Context, id string) (bool, error) {
	if err := s.UserRepository.DeleteByID(id); err != nil {
		return false, err
	}
	return true, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.UserRepository.GetByField("email", email)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (s *UserService) VerifyPassword(ctx context.Context, plainPasswordInput string, hashedPassword string) (bool, error)  {
	err := utils.ComparePassword(plainPasswordInput, hashedPassword)
	if err != nil {
		return false, nil
	}

	return true, nil
}