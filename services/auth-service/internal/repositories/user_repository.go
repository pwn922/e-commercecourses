package repositories

import (
	"github.com/pwn922/auth-service/graph/model"
	models "github.com/pwn922/auth-service/internal/models"
)

type UserRepository interface {
	Create(userInput *model.UserInput) (*models.User, error)
	GetByID(id int) (*models.User, error)
	List() ([]*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(userInput *model.UserInput, id int) (*models.User, error)
	DeleteByID(id int) error
}
