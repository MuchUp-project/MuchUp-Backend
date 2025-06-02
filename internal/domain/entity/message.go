package entity


import (
	"time"
	"errors"
	"github.com/google/uuid"
)

type Message struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMessage(userid,text string) (*Message,error) {
	if len(text) == 0 {
		return nil,errors.New("text is required")

	}
	if len(text) > 1000 {
		return nil, errors.New("text is too long")
	}
	message := &Message{
		ID : uuid.New().String(),
		UserID : userid,
		Text : text,
		CreatedAt : time.Now(),
	}
	return message , nil
}