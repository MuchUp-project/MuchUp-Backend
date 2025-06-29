package mocks
import (
	"MuchUp/backend/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)
type MockUserRepository struct {
	mock.Mock
}
func (m *MockUserRepository) CreateUser(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepository) GetUserByID(id string) (*entity.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}
func (m *MockUserRepository) UpdateUser(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepository) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockUserRepository) GetUsers(limit, offset int) ([]*entity.User, error) {
	args := m.Called(limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.User), args.Error(1)
}
func (m *MockUserRepository) GetUsersByGroup(groupID string) ([]*entity.User, error) {
	args := m.Called(groupID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.User), args.Error(1)
}
func (m *MockUserRepository) GetUserByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}
