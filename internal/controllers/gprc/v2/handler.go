// internal/controllers/grpc/v2/handler.go
package v2

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"MuchUp/backend/internal/domain/usecase"
	"MuchUp/backend/pkg/logger"
	pb "MuchUp/backend/proto/gen/go/v2"
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

// エラーハンドリングのヘルパー関数
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

// Health Check
func (h *GrpcHandler) HealthCheck(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	// This is just a dummy implementation to satisfy a potential health check.
	// We reuse GetUserRequest and User for simplicity.
	return &pb.User{Id: "health-ok"}, nil
}
