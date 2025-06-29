package message
import (
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repository"
	"MuchUp/backend/internal/domain/usecase"
	"time"
	"errors"
)
type messageUsecase struct {
	messageRepo repository.MessageRepository
	userRepo repository.UserRepository
}
func NewMessageUsecase(messageRepo repository.MessageRepository,userRepo repository.UserRepository) usecase.MessageUsecase {
	return &messageUsecase{
		messageRepo: messageRepo,
	    userRepo:userRepo,
	}
}
func (u *messageUsecase) CreateMessage(message *entity.Message) (*entity.Message, error) {
	err := u.messageRepo.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (u *messageUsecase) GetMessage(id string) (*entity.Message, error) {
    return u.messageRepo.GetMessageByID(id)
}


func (u *messageUsecase) GetMessageByID(id string) (*entity.Message, error) {
	return u.messageRepo.GetMessageByID(id)
}

func (u *messageUsecase)  SendMessage(message *entity.Message) error {
	if message.SenderID == "" {
		return errors.New("user id is required")
	}
	user, err := u.userRepo.GetUserByID(message.SenderID)
	if err != nil {
		return err
	}
	if user.IsBlockedUsers[message.SenderID] {
		return errors.New("user is blocked")
	}
	message.CreatedAt = time.Now()
	return u.messageRepo.CreateMessage(message)
}

func (u *messageUsecase) UnSentMessage(message *entity.Message) error {
	return u.UnSentMessage(message)
}



func (u *messageUsecase) UpdateMessage(message *entity.Message) (*entity.Message, error) {
	err := u.messageRepo.UpdateMessage(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
func (u *messageUsecase) DeleteMessage(id string) error {
	return u.messageRepo.DeleteMessage(id)
}
func (u *messageUsecase) GetMessagesByGroup(groupID string, limit, offset int) ([]*entity.Message, error) {
	return nil, nil
}
