package schema
import (
	"time"
	"gorm.io/gorm"
)
type MessageSchema struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Text      string `gorm:"type:text;not null"`
	SenderID  *string
	GroupID   string          `gorm:"type:uuid"`
	Group     ChatGroupSchema `gorm:"foreignKey:GroupID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
func (MessageSchema) TableName() string {
	return "messages"
}
