package schema

import (
	"gorm.io/gorm"
	"time"
)


type MessageSchema struct {
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Text string `gorm:"not null"`

	MaxLength int `gorm:"not null"`

	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	GroupID string `gorm:"not null"`
	ChatGroup GroupSchema  `gorm:"foreignKey:GroupID"`

	SenderID *string 
	Sender *UserSchema 

	IsDeleted bool `gorm:"default:false"`
	IsRead bool `gorm:"default:false"`
	IsDeletedBySender bool `gorm:"default:false"`
	IsDeletedBySystem bool `gorm:"default:false"`

	GroupMemnerShips GroupMembershipSchema `gorm:"foreignKey:ChatRoomID"`
	
}