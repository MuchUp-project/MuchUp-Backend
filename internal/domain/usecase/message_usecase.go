package usecase

import (
	"MuchUp/backend/internal/domain/entity"

)



type MessageUsecase interface {
	SendMessage(message *entity.Message) error
	UnSentMessage(message *entity.Message) error
	GetMessage(id string) (*entity.Message,error)
	GetMessageByID(id string) (*entity.Message,error)
	CreateMessage(message *entity.Message) (*entity.Message,error)
	UpdateMessage(message *entity.Message) (*entity.Message,error)
	DeleteMessage(id string) error
	GetMessagesByGroup(groupID string,limit,offset int) ([]*entity.Message,error)
}



