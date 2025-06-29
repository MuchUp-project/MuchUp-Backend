package repositories
import (
	"gorm.io/gorm"
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repository"
	"MuchUp/backend/internal/infrastructure/database/mapper"
	"MuchUp/backend/internal/infrastructure/database/schema"
)
type messageRepository struct {
	db *gorm.DB
}
func NewMessageRepository(db *gorm.DB) repository.MessageRepository {
	return &messageRepository{db: db}
}
func (r *messageRepository) CreateMessage(message *entity.Message) error {
	msgSchema := mapper.ToMessageSchema(message)
	return r.db.Create(&msgSchema).Error
}
func (r *messageRepository) GetMessageByID(id string) (*entity.Message, error) {
	var msgSchema schema.MessageSchema
	if err := r.db.First(&msgSchema, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return mapper.ToMessageEntity(&msgSchema), nil
}
func (r *messageRepository) GetMessagesByUserID(userID string) ([]*entity.Message, error) {
	return nil, nil
}
func (r *messageRepository) UpdateMessage(message *entity.Message) error {
	msgSchema := mapper.ToMessageSchema(message)
	return r.db.Save(&msgSchema).Error
}
func (r *messageRepository) DeleteMessage(id string) error {
	return r.db.Delete(&schema.MessageSchema{}, "id = ?", id).Error
}
func (r *messageRepository) GetMessagesByGroup(groupID string, limit, offset int) ([]*entity.Message, error) {
	var messagesSchema []schema.MessageSchema
	err := r.db.Where("group_id = ?", groupID).Limit(limit).Offset(offset).Order("created_at DESC").Find(&messagesSchema).Error
	if err != nil {
		return nil, err
	}
	messages := make([]*entity.Message, len(messagesSchema))
	for i, msgSchema := range messagesSchema {
		messages[i] = mapper.ToMessageEntity(&msgSchema)
	}
	return messages, nil
}
