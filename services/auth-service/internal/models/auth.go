package models

type LoginUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterUserInput struct {
	FirstName  string `json:"first_name" validate:"required"`
	MiddleName string `json:"middle_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	Role       string `json:"role" validate:"required"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
}