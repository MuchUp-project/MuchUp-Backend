package repositories

import (
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repository"
	"MuchUp/backend/internal/infrastructure/database/mapper"

	"gorm.io/gorm"
)
type userRepository struct {
	db *gorm.DB
}
func (r *userRepository) CreateUser(user *entity.User) error {
	userShema := mapper.ToUserSchema(user)
	err := r.db.Create(userShema).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetUserByPhone(phone string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}
func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&entity.User{}, "id = ?", id).Error
}
func (r *userRepository) GetUsers(limit, offset int) ([]*entity.User, error) {
	return nil, nil
}
func (r *userRepository) GetUsersByGroup(groupID string) ([]*entity.User, error) {
	return nil, nil
}
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}
