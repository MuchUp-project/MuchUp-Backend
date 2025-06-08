package schema

import (
	"time"

	"gorm.io/gorm"
)

type ChatGroupSchema struct {
	ID        string       `gorm:"type:uuid;primaryKey"`
	Name      string       `gorm:"type:varchar(100);not null"`
	Users     []UserSchema `gorm:"many2many:user_chat_groups;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ChatGroupSchema) TableName() string {
	return "chat_groups"
}
