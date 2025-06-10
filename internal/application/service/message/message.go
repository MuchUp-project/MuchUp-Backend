package message
import (
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repositories"
	"MuchUp/backend/internal/domain/usecase"
)
type messageUsecase struct {
	messageRepo repositories.MessageRepository
}
func NewMessageUsecase(messageRepo repositories.MessageRepository) usecase.MessageUsecase {
	return &messageUsecase{messageRepo: messageRepo}
}
func (u *messageUsecase) CreateMessage(message *entity.Message) (*entity.Message, error) {
	err := u.messageRepo.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
func (u *messageUsecase) GetMessageByID(id string) (*entity.Message, error) {
	return u.messageRepo.GetMessageByID(id)
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
