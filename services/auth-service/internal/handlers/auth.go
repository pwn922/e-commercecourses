package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pwn922/auth-service/internal/models"
	"github.com/pwn922/auth-service/internal/services/auth"
)

type AuthHandlers struct {
	AuthService *auth.AuthService
	Validator   *validator.Validate
}

func NewAuthHandlers(authService *auth.AuthService) *AuthHandlers {
	return &AuthHandlers{
		AuthService: authService,
		Validator:   validator.New(),
	}
}

func (h *AuthHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginInput models.LoginUserInput

	if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
		http.Error(w, "Invalid input: failed to parse JSON", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(loginInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authResponse, err := h.AuthService.Login(context.Background(), &loginInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(authResponse)
}

func (h *AuthHandlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userInput models.RegisterUserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid input: failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Validar el input
	if err := h.Validator.Struct(userInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authResponse, err := h.AuthService.Register(context.Background(), &userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(authResponse)
}
