package graph

import (
	"github.com/pwn922/auth-service/internal/repositories"
	"github.com/pwn922/auth-service/internal/services/auth"
	"github.com/pwn922/auth-service/internal/services/jwt"
	"github.com/pwn922/auth-service/internal/services/user"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    AuthService *auth.AuthService
}

func NewResolver(db *gorm.DB) *Resolver {
    userRepository := repositories.NewUserRepository(db)
    userService := user.NewUserService(userRepository)
    jwtService := jwt.NewJWTService()
	authService := auth.NewAuthService(userService, jwtService)
    return &Resolver{
        AuthService: authService,
    }
}