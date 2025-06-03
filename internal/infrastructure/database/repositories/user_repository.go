package repositories

import (
	
	"gorm.io/gorm"
	"github.com/MuchUp/backend/internal/domain/repositories"
	"github.com/MuchUp/backend/internal/domain/entity"
)


type userRepository struct {
	db *gorm.DB
}

	// GetUserByEmail(email string) (*entity.User,error)
	// GetUserByPhone(phone string) (*entity.User,error)
	// GetUserByID(id string) (*entity.User,error)
func (r *userRepository) CreateUser(user *entity.User) error {
	err :=  r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User,error) {
	var user entity.User
	err := r.db.Where("email = ?",email).First(&user).Error
	if err != nil {
		return nil,err
	}
	return &user,nil
}

func (r *userRepository) GetUserByPhone(phone string) (*entity.User,error) {
	var user entity.User
	err := r.db.Where("phone = ?",phone).First(&user).Error
	if err != nil {
		return nil,err
	}
	return &user,nil
}

func (r *userRepository) GetUserByID(id string) (*entity.User,error) {
	var user entity.User
	err := r.db.Where("id = ?",id).First(&user).Error
	if err != nil {
		return nil,err
	}
	return &user,nil
}


func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{db:db}
}