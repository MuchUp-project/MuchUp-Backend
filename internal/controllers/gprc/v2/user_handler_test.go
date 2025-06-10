package v2
import (
	"context"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/usecase"
	"MuchUp/backend/pkg/logger"
	pb "MuchUp/backend/proto/gen/go/v2"
)
func TestGrpcHandler_CreateUser(t *testing.T) {
	mockUserUsecase := new(usecase.MockUserUsecase)
	mockLogger := logger.New("error")
	handler := NewGrpcHandler(mockUserUsecase, nil, mockLogger)
	email := "test@example.com"
	testUser := &entity.User{
		ID:         "user-123",
		NickName:   "testuser",
		Email:      &email,
		AuthMethod: entity.AuthMethodEmail,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	mockUserUsecase.On("CreateUser", mock.AnythingOfType("*entity.User")).Return(testUser, nil)
	req := &pb.CreateUserRequest{
		NickName:     "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	res, err := handler.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testUser.ID, res.Id)
	assert.Equal(t, testUser.NickName, res.NickName)
	assert.Equal(t, *testUser.Email, res.Email)
	mockUserUsecase.AssertExpectations(t)
}
func TestGrpcHandler_GetUser(t *testing.T) {
	mockUserUsecase := new(usecase.MockUserUsecase)
	mockLogger := logger.New("error")
	handler := NewGrpcHandler(mockUserUsecase, nil, mockLogger)
	email := "test@example.com"
	testUser := &entity.User{
		ID:         "user-123",
		NickName:   "testuser",
		Email:      &email,
		AuthMethod: entity.AuthMethodEmail,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	mockUserUsecase.On("GetUserByID", "user-123").Return(testUser, nil)
	req := &pb.GetUserRequest{Id: "user-123"}
	res, err := handler.GetUser(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, testUser.ID, res.Id)
	mockUserUsecase.AssertExpectations(t)
}
func TestGrpcHandler_GetUser_NotFound(t *testing.T) {
	mockUserUsecase := new(usecase.MockUserUsecase)
	mockLogger := logger.New("error")
	handler := NewGrpcHandler(mockUserUsecase, nil, mockLogger)
	mockUserUsecase.On("GetUserByID", "non-existent-id").Return(nil, usecase.ErrNotFound)
	req := &pb.GetUserRequest{Id: "non-existent-id"}
	res, err := handler.GetUser(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, res)
	mockUserUsecase.AssertExpectations(t)
}
