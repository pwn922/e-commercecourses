// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Query struct {
}

type Role struct {
	ID          string  `json:"id"`
	RoleName    string  `json:"roleName"`
	Description *string `json:"description,omitempty"`
	Users       []*User `json:"users"`
}

type RoleInput struct {
	RoleName    string  `json:"role_name"`
	Description *string `json:"description,omitempty"`
}

type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Role       *Role  `json:"role"`
}

type UserInput struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}
