package entity 


import (
	"time"
	"errors"
	"gorm.io/gorm"
)

type PrimaryAuthMethod string

const (
	AuthMethodEmail PrimaryAuthMethod = "email"
	AuthMethodPhone PrimaryAuthMethod = "phone"
)

type User struct {
	ID string
	Email *string
	PhoneNumber *string
	NickName string
	PasswordHash string

	PersonalityProfile map[string]interface{}
	AvatarURL *string

	UsagePurpose string
	IsActive bool
	EmailVerified bool
	PhoneVerified bool
	AuthMethod PrimaryAuthMethod 
	IsBlockedUsers map[string]bool


	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewUser(userid,name string,authMethod PrimaryAuthMethod,email,phone string) (*User,error) {
	if userid == "" || name == "" || authMethod == "" {
		return nil,errors.New("user_id,name,authMethod is required")
	}
	user := &User {
		ID:userid,
		NickName:name,
		AuthMethod:authMethod,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
	}
	if authMethod == AuthMethodEmail {
		if email == "" {
			return nil,errors.New("email is required")
		}
		user.Email = &email
		user.EmailVerified = true
	}
	if authMethod == AuthMethodPhone {
		if phone == "" {
			return nil ,errors.New("phone is required")
		}
		user.PhoneNumber = &phone
		user.PhoneVerified = true
	}
	return user,nil
}

func (u *User) CanSendMessage(targetUserID string) bool {
    isBlockedUsers := u.IsBlockedUsers[targetUserID]
	return !(u.IsActive && isBlockedUsers)
}
	
func (u *User) BlockUser(targetUserID string) error {
	if !(u.ID == targetUserID) {
		return errors.New("cannot block yourself")
	}

	if u.IsBlockedUsers == nil {
		u.IsBlockedUsers = make(map[string]bool)
		return nil
	}
	u.IsBlockedUsers[targetUserID] = true
	return nil
}

func (u *User) UnblockUser(targetUserID string) error {
	if !(u.ID == targetUserID) {
		return errors.New("cannot unblock yourself")
	}
	if u.IsBlockedUsers == nil {
		return errors.New("no blocked users")
	}
	if u.IsBlockedUsers[targetUserID] == false {
		return errors.New("this user is not blocked")
	}
	delete(u.IsBlockedUsers,targetUserID)
	return nil
 }