package usecase
import "MuchUp/backend/internal/domain/entity"
type MessageService interface {
	SendMessage(senderID, groupID string,content string) error
	UnSentMessage(message *entity.Message) error
	GetMessage(message *entity.Message) (*entity.Message,error)
}