package usecase 
import (
	"MuchUp/backend/internal/domain/entity"
)



type UserUsecase interface {
	GetUserByID(id string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User,error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id string) error
	GetUsers(limit, offset int) ([]*entity.User, error)
	Login(email, password string) (string, error)
	JoinGroup(userID, groupID string) error
	LeaveGroup(userID, groupID string) error
	GetUsersByGroup(groupID string) ([]*entity.User, error)
}