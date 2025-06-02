package repositories

import (
	"githib.com/MuchUp/backend/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User,error)
	GetUserByPhone(phone string) (*entity.User,error)
	GetUserByID(id string) (*entity.User,error)
}