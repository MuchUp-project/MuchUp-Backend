package entity

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type ChatRoom struct {
	ID         string    `json:"id"`
	Members    []User    `json:"users"`
	Messages   []Message `json:"messages"`
	MaxMembers int
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewChatRoom(users []User) (*ChatRoom, error) {
	maxMember := 6
	if len(users) > maxMember {
		return nil, errors.New("max 6 users are allowed")
	}
	return &ChatRoom{
		ID:         uuid.New().String(),
		Members:    users,
		Messages:   []Message{},
		MaxMembers: maxMember,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil

}

func (c *ChatRoom)IsMember(user User) bool {
	for _,member := range c.Members {
		if member.UserID == user.UserID {
			return true
		}
	
	}
	return false
}

func (c *ChatRoom) AddMessageLog(user User, message Message) error {
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

