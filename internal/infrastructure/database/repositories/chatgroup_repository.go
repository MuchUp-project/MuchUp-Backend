package repositories
import (
	"errors"
	"fmt"
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repository"
	"MuchUp/backend/internal/domain/usecase"
	"MuchUp/backend/internal/infrastructure/database/mapper"
	"MuchUp/backend/internal/infrastructure/database/schema"
	"gorm.io/gorm"
)
type chatGroupRepository struct {
	db *gorm.DB
}
func NewChatGroupRepository(db *gorm.DB) repository.ChatGroupRepository {
	return &chatGroupRepository{db: db}
}
func (r *chatGroupRepository) CreateGroup(group *entity.ChatGroup) (*entity.ChatGroup, error) {
	groupSchema := mapper.ToGroupSchema(group)
	if err := r.db.Create(groupSchema).Error; err != nil {
		return nil, err
	}
	return r.GetGroupByID(group.ID)
}
func (r *chatGroupRepository) GetGroupByID(id string) (*entity.ChatGroup, error) {
	var groupSchema schema.ChatGroupSchema
	err := r.db.Preload("Users").First(&groupSchema, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("group not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get group: %w", err)
	}
	return mapper.ToGroupEntity(&groupSchema), nil
}
func (r *chatGroupRepository) FindGroupWithAvailableSlots() (*entity.ChatGroup, error) {
	var groupSchema schema.ChatGroupSchema
	err := r.db.Preload("Users").
		Joins("LEFT JOIN user_chat_groups on user_chat_groups.chat_group_id = chat_groups.id").
		Group("chat_groups.id").
		Having("COUNT(user_chat_groups.user_id) < 6").
		Order("COUNT(user_chat_groups.user_id) DESC").
		First(&groupSchema).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, usecase.ErrNotFound
		}
		return nil, err
	}
	return mapper.ToGroupEntity(&groupSchema), nil
}
func (r *chatGroupRepository) AddUserToGroup(userID, groupID string) error {
	user := schema.UserSchema{ID: userID}
	group := schema.ChatGroupSchema{ID: groupID}
	return r.db.Model(&group).Association("Users").Append(&user)
}
