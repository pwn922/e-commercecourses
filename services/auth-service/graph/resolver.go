package graph

import (
	"github.com/pwn922/auth-service/internal/services/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    AuthService *auth.AuthService
}

func NewResolver(authService *auth.AuthService) *Resolver {
    return &Resolver{
        AuthService: authService,
    }
}