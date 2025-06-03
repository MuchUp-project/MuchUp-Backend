package usecase 

import (
	"github.com/MuchUp/backend/internal/domain/entity"
)

type UserAuthService interface {
	RegisterUser(user *entity.User) error
}