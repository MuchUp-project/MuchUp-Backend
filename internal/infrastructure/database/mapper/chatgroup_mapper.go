package mapper

import (
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/infrastructure/database/schema"
)

func ToGroupSchema(group *entity.ChatGroup) *schema.ChatGroupSchema {
	users := make([]schema.UserSchema, len(group.Members))
	for i, member := range group.Members {
		users[i] = *ToUserSchema(&member)
	}

	return &schema.ChatGroupSchema{
		ID:    group.ID,
		Name:  group.Name,
		Users: users,
	}
}

func ToGroupEntity(groupSchema *schema.ChatGroupSchema) *entity.ChatGroup {
	if groupSchema == nil {
		return nil
	}
	members := make([]entity.User, len(groupSchema.Users))
	for i, userSchema := range groupSchema.Users {
		members[i] = *ToUserEntity(&userSchema)
	}

	return &entity.ChatGroup{
		ID:        groupSchema.ID,
		Name:      groupSchema.Name,
		Members:   members,
		CreatedAt: groupSchema.CreatedAt,
		UpdatedAt: groupSchema.UpdatedAt,
	}
}
