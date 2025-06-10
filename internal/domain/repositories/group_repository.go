package repositories 
import (
	"MuchUp/backend/internal/domain/entity"
)
type GroupRepository interface {
	CreateGroup(group *entity.ChatGroup) error
	DeleteGroup(groupID string) error
	GetGroupByID(groupID string) (*entity.ChatGroup,error)
	GetGroupByUserID(userID string) ([]*entity.ChatGroup,error)		
}