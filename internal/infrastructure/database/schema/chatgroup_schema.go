package schema 


import (
	"time"
	"gorm.io/gorm"
)


type GroupSchema struct {
	Group string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Messages []MessageSchema `gorm:"foreignKey:GroupID"`
	Memberships []GroupMembershipSchema `gorm:"foreignKey:GroupID"`
}

