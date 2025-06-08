package user

import (
	"errors"
	"time"

	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repositories"
	// "MuchUp/backend/internal/domain/usecase"
)

type sendMessageService struct {
	messageRepo repositories.MessageRepository
	userRepo    repositories.UserRepository
}

// func NewSendMessageService(messageRepo repositories.MessageRepository, userRepo repositories.UserRepository) usecase.SendMessageService {
// 	return &sendMessageService{messageRepo: messageRepo, userRepo: userRepo}
// }

func (s *sendMessageService) SendMessage(message *entity.Message) error {
	if message.SenderID == "" {
		return errors.New("user id is required")
	}
	user, err := s.userRepo.GetUserByID(message.SenderID)
	if err != nil {
		return err
	}
	if user.IsBlockedUsers[message.SenderID] {
		return errors.New("user is blocked")
	}
	message.CreatedAt = time.Now()
	return s.messageRepo.CreateMessage(message)
}
