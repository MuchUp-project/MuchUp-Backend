package v2
import (
	"context"
	"MuchUp/backend/internal/domain/entity"
	pb "MuchUp/backend/proto/gen/go/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)
func (h *GrpcHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := &entity.User{
		NickName:     req.GetNickName(),
		Email:        &req.Email,
		PasswordHash: req.Password,
	}
	createdUser, err := h.userUsecase.CreateUser(user)
	if err != nil {
		return nil, h.handleError(ctx, "CreateUser", err)
	}
	return &pb.User{
		Id:        createdUser.ID,
		NickName:  createdUser.NickName,
		Email:     *createdUser.Email,
		CreatedAt: timestamppb.New(createdUser.CreatedAt),
		UpdatedAt: timestamppb.New(createdUser.UpdatedAt),
	}, nil
}
func (h *GrpcHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := h.userUsecase.GetUserByID(req.GetId())
	if err != nil {
		return nil, h.handleError(ctx, "GetUser", err)
	}
	return &pb.User{
		Id:        user.ID,
		NickName:  user.NickName,
		Email:     *user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}
