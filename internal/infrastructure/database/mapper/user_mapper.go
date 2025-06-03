 package mapper

 import (
	"github.com/MuchUp/backend/internal/domain/entity"
	"github.com/MuchUp/backend/internal/infrastructure/database/schema"
 )

 func ToUserSchema(user *entity.User)  *schema.UserSchema {
		userSchema :=  &schema.UserSchema{
			ID: user.ID,
			NickName: user.NickName,
			PasswordHash: user.PasswordHash,
			EmailVerified: user.EmailVerified,
			UsagePurpose: user.UsagePurpose,
			IsActive: user.IsActive,
			AvatarURL: user.AvatarURL,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
			PersonalityProfile: user.PersonalityProfile,
			PrimaryAuthMethod: string(user.AuthMethod),
		}	
		switch user.AuthMethod {
			case entity.AuthMethodEmail :{
				userSchema.Email = user.Email
				userSchema.EmailVerified = true
			}
			case entity.AuthMethodPhone : {
				userSchema.PhoneNumber= user.PhoneNumber
				userSchema.PhoneVerified = true
			}
		}
		return userSchema
	}
	