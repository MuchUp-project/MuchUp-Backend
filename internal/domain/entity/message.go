package entity

import (
	"errors"
	"time"

	"MuchUp/backend/utils"
)

type Message struct {
	MessageID string  `json:"message_id"`
	SenderID  string  `json:"user_id"`
	GroupID   string  `json:"group_id"`
	Text      *string `json:"text"`
	Image     *string `json:"image"`
	Video     *string `json:"video"`
	Sticker   *string `json:"sticker"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewMessage(userid, groupid, text string) (*Message, error) {
	if len(text) == 0 {
		return nil, errors.New("text is required")

	}
	if len(text) > 1000 {
		return nil, errors.New("text is too long")
	}
	message := &Message{
		MessageID: utils.GenerateUUID(),
		SenderID:  userid,
		GroupID:   groupid,
		Text:      &text,

		CreatedAt: time.Now(),
	}
	return message, nil
}

func (m *Message) CanSendMessage(senderID string) bool {
	if m.SenderID == senderID {
		if m.Text == nil && m.Image == nil && m.Video == nil && m.Sticker == nil {
			return false
		}
		if m.Text != nil || len(*m.Text) > 1000 {
			return false
		}
		return true
	}
	return false
}
