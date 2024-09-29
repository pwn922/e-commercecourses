package user

import (
	"context"

	"github.com/pwn922/auth-service/graph/model"
	"github.com/pwn922/auth-service/internal/models"
	"github.com/pwn922/auth-service/internal/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) GetUser(ctx context.Context, id int) (*models.User, error) {
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

func (s *UserService) CreateUser(ctx context.Context, newUserInput *model.UserInput) (*models.User, error) {
	createdUser, err := s.UserRepository.Create(newUserInput)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) UpdateUser(ctx context.Context, updateUserInput *model.UserInput, id int) (*models.User, error) {
	updatedUser, err := s.UserRepository.Update(updateUserInput, id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) (bool, error) {
	if err := s.UserRepository.DeleteByID(id); err != nil {
		return false, err
	}
	return true, nil
}
