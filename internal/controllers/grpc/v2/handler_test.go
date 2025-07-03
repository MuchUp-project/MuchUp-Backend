package v2

import (
	
	"context"
	"errors"
	"testing"
	"MuchUp/backend/pkg/logger"
	"MuchUp/backend/internal/domain/entity"
	mock_usecase "MuchUp/backend/internal/domain/usecase/mocks"
	pd "MuchUp/backend/proto/gen/go/v2"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

)

func TestHandler_CreateUser(t *testing.T) {

	ctx := context.Background()
	assert := require.New(t)
    logger :=logger.NewLogger()

	req := &pd.CreateUserRequest{
		NickName: "test",
		Email: "test@example.com",
		Password: "testpasswordhash",
		}

	createdUser := &entity.User{
		NickName: req.NickName,
		Email: &req.Email,

	}

	t.Run("success pattern",func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)

		mockUserUsecase.EXPECT().CreateUser(gomock.Any()).Return(createdUser, nil)

		handler := NewGrpcHandler(mockUserUsecase, nil, logger)

		res,err:= handler.CreateUser(ctx, req)

		assert.NoError(err)
		assert.NotNil(res)
		assert.Equal(res.NickName, createdUser.NickName)
		assert.Equal(res.Email, *createdUser.Email)


	})
	t.Run(("error pattern"),func(t *testing.T){
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)
		expectedErr := errors.New("email already exists")


		mockUserUsecase.EXPECT().CreateUser(gomock.Any()).Return(nil,expectedErr)

		handler := NewGrpcHandler(mockUserUsecase, nil, logger)

		res,err := handler.CreateUser(ctx, req)

		assert.Error(err)
		assert.Nil(res)
	})
}

func TestHandler_GetUser(t *testing.T) {
	ctx := context.Background()
	assert := require.New(t)
	logger :=logger.NewLogger()

	testID := "testID"
	testEmail := "test@example.com"
	testName := "testName"

	req := &pd.GetUserRequest{
		Id: testID,
	}

	gettingUser := &entity.User{
		ID:testID,
		Email: &testEmail,
		NickName: testName,
	}


	t.Run("success pattern",func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)

		mockUserUsecase.EXPECT().GetUserByID(testID).Return(gettingUser,nil)

		handler := NewGrpcHandler(mockUserUsecase,nil,logger)
		res,err := handler.GetUser(ctx,req)

		assert.NoError(err)
		assert.NotNil(res)
		assert.Equal(testID,res.Id)
		assert.Equal(testEmail,res.Email)
	})

	t.Run("error pattern",func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)

		mockUserUsecase.EXPECT().GetUserByID(gomock.Any()).Return(nil,errors.New("error"))


		handler := NewGrpcHandler(mockUserUsecase,nil,logger)
		res,err := handler.GetUser(ctx,req)

		assert.Error(err)
		assert.Nil(res)

	})




}
