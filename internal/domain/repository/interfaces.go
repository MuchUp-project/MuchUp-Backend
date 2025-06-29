package repository
import "MuchUp/backend/internal/domain/entity"



type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByID(id string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id string) error
	GetUsers(limit, offset int) ([]*entity.User, error)
	GetUsersByGroup(groupID string) ([]*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
type MessageRepository interface {
	CreateMessage(message *entity.Message) error
	GetMessageByID(id string) (*entity.Message, error)
	GetMessagesByUserID(userID string) ([]*entity.Message, error)
	UpdateMessage(message *entity.Message) error
	DeleteMessage(id string) error
	GetMessagesByGroup(groupID string, limit, offset int) ([]*entity.Message, error)
}
type ChatGroupRepository interface {
	CreateGroup(group *entity.ChatGroup) (*entity.ChatGroup, error)
	GetGroupByID(id string) (*entity.ChatGroup, error)
	AddUserToGroup(userID, groupID string) error
	FindGroupWithAvailableSlots() (*entity.ChatGroup, error)
}
