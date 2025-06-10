package entity
import (
	"time"
	"github.com/google/uuid"
)
type ChatGroup struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Members    []User   `json:"members"`
	MessageIDs []string `json:"messages"`
	MaxMembers int
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
func NewChatGroup(defaultName string, maxMember int, initialUser User) *ChatGroup {
	return &ChatGroup{
		ID:         uuid.New().String(),
		Name:       defaultName,
		Members:    []User{initialUser},
		MaxMembers: maxMember,
		CreatedAt:  time.Now(),
	}
}
func (c *ChatGroup) IsMember(userID string) bool {
	for _, member := range c.Members {
		if userID == member.ID {
			return true
		}
	}
	return false
}
func (c *ChatGroup) CanAddMemnber() bool {
	if len(c.Members) >= c.MaxMembers {
		return false
	}
	return true
}
func (c *ChatGroup) CanDeleteGroup() bool {
	if len(c.Members) > 1 {
		return false
	}
	return true
}
