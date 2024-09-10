package graph

import "github.com/pwn922/users-service/internal/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    UserService *services.UserService
}