package user
import (
	"time"
	"errors"
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/repositories"
	"MuchUp/backend/internal/domain/usecase"
)
type userAuthService struct {
	userRepo repositories.UserRepository
}
func NewUserAuthService(userRepo repositories.UserRepository) usecase.UserAuthService {
	return &userAuthService{userRepo:userRepo}
}
func (s *userAuthService) RegisterUser(user *entity.User) error {
	if !user.IsActive  {
		return errors.New("user is not acitve")
	}
	user.CreatedAt = time.Now()
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
