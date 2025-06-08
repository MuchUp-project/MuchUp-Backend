package usecase

import (
	"errors"

	"MuchUp/backend/internal/domain/entity"
)

// エラー定義
var (
	ErrNotFound         = errors.New("not found")
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrPermissionDenied = errors.New("permission denied")
)

// UserUsecase ユーザー関連のユースケースインターフェース
type UserUsecase interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id string) error
	GetUsers(limit, offset int) ([]*entity.User, error)
	Login(email, password string) (string, error)
	JoinGroup(userID, groupID string) error
	LeaveGroup(userID, groupID string) error
	GetUsersByGroup(groupID string) ([]*entity.User, error)
}

// MessageUsecase メッセージ関連のユースケースインターフェース
type MessageUsecase interface {
	CreateMessage(message *entity.Message) (*entity.Message, error)
	GetMessageByID(id string) (*entity.Message, error)
	UpdateMessage(message *entity.Message) (*entity.Message, error)
	DeleteMessage(id string) error
	GetMessagesByGroup(groupID string, limit, offset int) ([]*entity.Message, error)
}

type GroupUsecase interface {
	FindOrCreateGroupForUser(user *entity.User) (*entity.ChatGroup, error)
	AddUserToGroup(userID, groupID string) error
}
