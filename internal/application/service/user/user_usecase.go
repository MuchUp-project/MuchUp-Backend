package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repositories"
	"MuchUp/backend/internal/domain/usecase"
)

type userUsecase struct {
	userRepo  repositories.UserRepository
	groupRepo repositories.ChatGroupRepository
	groupUc   usecase.GroupUsecase
}

func NewUserUsecase(userRepo repositories.UserRepository, groupUc usecase.GroupUsecase) usecase.UserUsecase {
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

	// Find or create a group for the new user.
	// This operation is critical for the user flow but if it fails,
	// we might decide not to fail the whole user creation.
	// For now, we return an error. A more robust implementation could
	// add the user to a queue for later processing.
	_, err = u.groupUc.FindOrCreateGroupForUser(user)
	if err != nil {
		// Log the error but consider if we should return it
		// logging is handled within the usecase, so just return
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetUserByID(id string) (*entity.User, error) {
	return u.userRepo.GetUserByID(id)
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

	// TODO: Implement JWT token generation
	token := "dummy-jwt-token"

	return token, nil
}

// --- Placeholder Implementations ---

func (u *userUsecase) UpdateUser(user *entity.User) (*entity.User, error) {
	// TODO: Implement password update logic if password is provided
	// For now, assuming it's handled elsewhere or not part of this update
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
