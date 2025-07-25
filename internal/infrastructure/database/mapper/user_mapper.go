package mapper
import (
	"encoding/json"
	"log"
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/infrastructure/database/schema"
)
func ToUserSchema(user *entity.User) *schema.UserSchema {
	var profileJSON json.RawMessage
	if user.PersonalityProfile != nil {
		profileData, err := json.Marshal(user.PersonalityProfile)
		if err != nil {
			log.Printf("Error marshalling user profile: %v", err)
			profileJSON = json.RawMessage("{}")
		} else {
			profileJSON = profileData
		}
	}
	return &schema.UserSchema{
		ID:                 user.ID,
		NickName:           user.NickName,
		Email:              user.Email,
		PasswordHash:       user.PasswordHash,
		UsagePurpose:       user.UsagePurpose,
		PersonalityProfile: profileJSON,
		CreatedAt:          user.CreatedAt,
		UpdatedAt:          user.UpdatedAt,
	}
}
func ToUserEntity(userSchema *schema.UserSchema) *entity.User {
	var profile map[string]interface{}
	if userSchema.PersonalityProfile != nil {
		_ = json.Unmarshal(userSchema.PersonalityProfile, &profile)
	}
	return &entity.User{
		ID:                 userSchema.ID,
		NickName:           userSchema.NickName,
		Email:              userSchema.Email,
		PasswordHash:       userSchema.PasswordHash,
		UsagePurpose:       userSchema.UsagePurpose,
		PersonalityProfile: profile,
		CreatedAt:          userSchema.CreatedAt,
		UpdatedAt:          userSchema.UpdatedAt,
	}
}
