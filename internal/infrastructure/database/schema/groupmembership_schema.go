package schema

import (
	
	"time"
	"gorm.io/gorm"
)


type GroupMembershipSchema struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID string `gorm:"not null"`
	GroupID string `gorm:"not null"`

	JoinedAt time.Time `gorm:"default:now()"`
	LeftAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User UserSchema `gorm:"foreignKey:UserID"`
	Group GroupSchema `gorm:"foreignKey:GroupID"`

}

func (GroupMembershipSchema) TableName() string {
	return "group_membership"
}