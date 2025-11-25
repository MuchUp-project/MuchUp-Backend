package user

import (
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repository"
	"MuchUp/backend/internal/domain/usecase"
	"errors"

	"golang.org/x/crypto/bcrypt"
)
type userUsecase struct {
	userRepo  repository.UserRepository
	groupRepo repository.ChatGroupRepository
	groupUc   usecase.GroupUsecase
}
func NewUserUsecase(userRepo repository.UserRepository, groupUc usecase.GroupUsecase) usecase.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		groupUc:  groupUc,
	}
}
func (u *userUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hashedPassword)
	if err := u.userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	_, err = u.groupUc.FindOrCreateGroupForUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userUsecase) GetUserByID(id string) (*entity.User, error) {
	return u.userRepo.GetUserByID(id)
}
func (u *userUsecase) GetUserByEmail(email string) (*entity.User,error) {
	return u.userRepo.GetUserByEmail(email)
}
func (u *userUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", usecase.ErrNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	token := "dummy-jwt-token"
	return token, nil
}
func (u *userUsecase) UpdateUser(user *entity.User) (*entity.User, error) {

	return nil, errors.New("not implemented")
}
func (u *userUsecase) DeleteUser(id string) error {
	return errors.New("not implemented")
}
func (u *userUsecase) GetUsers(limit, offset int) ([]*entity.User, error) {
	return nil, errors.New("not implemented")
}
func (u *userUsecase) JoinGroup(userID, groupID string) error {
	return errors.New("not implemented")
}
func (u *userUsecase) LeaveGroup(userID, groupID string) error {
	return errors.New("not implemented")
}
func (u *userUsecase) GetUsersByGroup(groupID string) ([]*entity.User, error) {
	return nil, errors.New("not implemented")
}
