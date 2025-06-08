package usecase

import (
	"MuchUp/backend/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

// MockUserUsecase is a mock of UserUsecase interface
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserUsecase) GetUserByID(id string) (*entity.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserUsecase) UpdateUser(user *entity.User) (*entity.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserUsecase) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserUsecase) GetUsers(limit, offset int) ([]*entity.User, error) {
	args := m.Called(limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserUsecase) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) JoinGroup(userID, groupID string) error {
	args := m.Called(userID, groupID)
	return args.Error(0)
}

func (m *MockUserUsecase) LeaveGroup(userID, groupID string) error {
	args := m.Called(userID, groupID)
	return args.Error(0)
}

func (m *MockUserUsecase) GetUsersByGroup(groupID string) ([]*entity.User, error) {
	args := m.Called(groupID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.User), args.Error(1)
}

// MockMessageUsecase is a mock of MessageUsecase interface
type MockMessageUsecase struct {
	mock.Mock
}

func (m *MockMessageUsecase) CreateMessage(message *entity.Message) (*entity.Message, error) {
	args := m.Called(message)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Message), args.Error(1)
}

func (m *MockMessageUsecase) GetMessageByID(id string) (*entity.Message, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Message), args.Error(1)
}

func (m *MockMessageUsecase) UpdateMessage(message *entity.Message) (*entity.Message, error) {
	args := m.Called(message)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Message), args.Error(1)
}

func (m *MockMessageUsecase) DeleteMessage(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMessageUsecase) GetMessagesByGroup(groupID string, limit, offset int) ([]*entity.Message, error) {
	args := m.Called(groupID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Message), args.Error(1)
}
