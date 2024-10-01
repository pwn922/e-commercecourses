package auth

import (
	"context"

	"github.com/pwn922/auth-service/graph/model"
	"github.com/pwn922/auth-service/internal/models"
	"github.com/pwn922/auth-service/internal/services/jwt"
	"github.com/pwn922/auth-service/internal/services/user"
)

type ErrorResponse struct {
    Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
    return e.Message
}

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

func (a *AuthService) Register(ctx context.Context, registerUserInput *models.RegisterUserInput) (*models.AuthResponse, error) {
    user, err := a.UserService.CreateUser(ctx, registerUserInput)
    if err != nil || user == nil {
        return nil, err
    }

    token, err := a.JWTService.GenerateAccessToken(user.ID, user.Role.ID)
    if err != nil {
        return nil, err
    }

    return &models.AuthResponse{AccessToken: token}, nil
}

func (a *AuthService) Login(ctx context.Context, loginInput *models.LoginUserInput) (*models.AuthResponse, error) {
    const errorMessage = "Incorrect email or password."

    user, err := a.UserService.GetUserByEmail(ctx, loginInput.Email)
    if err != nil || user == nil {
        println("NO ENCONTRO EL CORREO");
        return nil, &ErrorResponse{Message: errorMessage}
    }

    valid, err := a.UserService.VerifyPassword(ctx, loginInput.Password, user.HashedPassword)
    if err != nil || !valid {
        println("LA CONTRASEÑA NO ES VALIDA");
        return nil, &ErrorResponse{Message: errorMessage}
    }

    token, err := a.JWTService.GenerateAccessToken(user.ID, user.Role.ID)
    if err != nil {
        return nil, err
    }

    return &models.AuthResponse{AccessToken: token}, nil
}

func (a *AuthService) GetProfileUser(ctx context.Context, id string) (*model.User, error) {
    user, err := a.UserService.GetUser(ctx, id)
    if err != nil || user == nil {
        return nil, err
    }

    return &model.User{
        ID:        user.ID,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Email:     user.Email,
        Role: &model.Role{
            ID:          user.Role.ID,
            RoleName:    user.Role.RoleName,
            Description: &user.Role.Description,
        },
    }, nil
}