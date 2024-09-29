package auth

import (
	"context"

	"github.com/pwn922/auth-service/graph/model"
	"github.com/pwn922/auth-service/internal/services/jwt"
	"github.com/pwn922/auth-service/internal/services/user"
)

type AuthService struct {
    UserService *user.UserService
    JWTService  *jwt.JWTService
}

func NewAuthService(userService *user.UserService, jwtService *jwt.JWTService) *AuthService {
    return &AuthService{
        UserService: userService,
        JWTService:  jwtService,
    }
}

func (a *AuthService) Login(ctx context.Context, userInput *model.UserInput) (*model.AuthResponse, error) {
    createdUser, err := a.UserService.CreateUser(ctx, userInput)
    if err != nil {
        return nil, err
    }

    token, err := a.JWTService.GenerateAccessToken(createdUser.ID, createdUser.Role.ID)
    if err != nil {
        return nil, err
    }
	
    return &model.AuthResponse{Token: token}, nil
}
