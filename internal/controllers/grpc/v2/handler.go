package v2

import (
	"MuchUp/backend/internal/domain/usecase"
	"MuchUp/backend/pkg/logger"
	pb "MuchUp/backend/proto/gen/go/v2"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcHandler struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedMessageServiceServer
	userUsecase    usecase.UserUsecase
	messageUsecase usecase.MessageUsecase
	logger         logger.Logger
}

func NewGrpcHandler(
	userUsecase usecase.UserUsecase,
	messageUsecase usecase.MessageUsecase,
	logger logger.Logger,
) *GrpcHandler {
	return &GrpcHandler{
		userUsecase:    userUsecase,
		messageUsecase: messageUsecase,
		logger:         logger,
	
	}
}
func (h *GrpcHandler) handleError(ctx context.Context, operation string, err error) error {
	if err == nil {
		return nil
	}
	h.logger.WithContext(ctx).WithError(err).Errorf("Failed to %s", operation)
	switch {
	case err == usecase.ErrNotFound:
		return status.Error(codes.NotFound, fmt.Sprintf("%s not found", operation))
	case err == usecase.ErrInvalidArgument:
		return status.Error(codes.InvalidArgument, err.Error())
	case err == usecase.ErrPermissionDenied:
		return status.Error(codes.PermissionDenied, "permission denied")
	default:
		return status.Error(codes.Internal, fmt.Sprintf("failed to %s", operation))
	}
}
func (h *GrpcHandler) HealthCheck(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{Id: "health-ok"}, nil
}

