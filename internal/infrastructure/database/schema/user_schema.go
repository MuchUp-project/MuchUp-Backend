package schema
import (
	"encoding/json"
	"errors"
	"time"
	"gorm.io/gorm"
)
type UserSchema struct {
	ID string `gorm:"type:uuid;primaryKey"`
	Email        *string `gorm:"type:varchar(255);unique"`
	PhoneNumber  *string `gorm:"size:20;uniqueIndex"`
	PasswordHash string  `gorm:"type:varchar(255);not null"`
	EmailVerified     bool   `gorm:"default:false"`
	PhoneVerified     bool   `gorm:"default:false"`
	PrimaryAuthMethod string `gorm:"size:10;not null"`
	EmailVerificationToken *string `gorm:"size:255"`
	PhoneVerificationCode  *string `gorm:"size:6"`
	VerificationCodeExpiry *time.Time
	NickName  string  `gorm:"type:varchar(50);not null"`
	AvatarURL *string `gorm:"size:500"`
	PersonalityProfile json.RawMessage `gorm:"type:jsonb"`
	UsagePurpose       string          `gorm:"type:varchar(255)"`
	IsActive             bool `gorm:"default:true"`
	NotificationsEnabled bool `gorm:"default:true"`
	TwoFactorEnabled     bool `gorm:"default:false"`
	LastLoginAt   *time.Time
	LoginAttempts int `gorm:"default:0"`
	LockedUntil   *time.Time
	ResetToken       *string `gorm:"size:255"`
	ResetTokenExpiry *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ChatGroups []ChatGroupSchema `gorm:"many2many:user_chat_groups;"`
	Messages   []MessageSchema   `gorm:"foreignKey:UserID"`
}
func (UserSchema) TableName() string {
	return "users"
}
func (u *UserSchema) BeforeCreate(tx *gorm.DB) error {
	if u.Email == nil && u.PhoneNumber == nil {
		return errors.New("either email or phone number must be provided")
	}
	return nil
}
