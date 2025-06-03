package entity

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type ChatGroup struct {
	ID         string    `json:"id"`
	Members    []User    `json:"users"`
	Messages   []Message `json:"messages"`
	MaxMembers int
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func NewChatRoom(users []User) (*ChatGroup, error) {
	maxMember := 6
	if len(users) > maxMember {
		return nil, errors.New("max 6 users are allowed")
	}
	return &ChatGroup{
		ID:         uuid.New().String(),
		Members:    users,
		Messages:   []Message{},
		MaxMembers: maxMember,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil

}

func (c *ChatGroup)IsMember(user User) bool {
	for _,member := range c.Members {
		if member.ID == user.ID {
			return true
		}
	
	}
	return false
}

func (c *ChatGroup) AddMessageLog(user User, message Message) error {
	if !c.IsMember(user) {
		return errors.New("this user is not a member of this chatroom")
	}
	if len(c.Messages) >= 1000 {
		return errors.New("message limit reached")
	}

	log.Println("recent message: %s", message.Text)
	c.Messages = append(c.Messages, message)
	c.UpdatedAt = time.Now()
	return nil
}

