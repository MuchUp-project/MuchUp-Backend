package v2
import (
	"context"
	"MuchUp/backend/internal/domain/entity"
	pb "MuchUp/backend/proto/gen/go/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)
func (h *GrpcHandler) CreateMessage(ctx context.Context, req *pb.CreateMessageRequest) (*pb.Message, error) {
	message, err := entity.NewMessage(req.GetSenderId(), req.GetGroupId(), req.GetText())
	if err != nil {
		return nil, h.handleError(ctx, "CreateMessage.NewMessage", err)
	}
	createdMessage, err := h.messageUsecase.CreateMessage(message)
	if err != nil {
		return nil, h.handleError(ctx, "CreateMessage.Create", err)
	}
	return &pb.Message{
		MessageId: createdMessage.MessageID,
		SenderId:  createdMessage.SenderID,
		GroupId:   createdMessage.GroupID,
		Text:      *createdMessage.Text,
		CreatedAt: timestamppb.New(createdMessage.CreatedAt),
	}, nil
}
func (h *GrpcHandler) GetMessagesByGroup(req *pb.GetMessagesByGroupRequest, stream pb.MessageService_GetMessagesByGroupServer) error {
	messages, err := h.messageUsecase.GetMessagesByGroup(req.GetGroupId(), int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return h.handleError(stream.Context(), "GetMessagesByGroup", err)
	}
	for _, msg := range messages {
		res := &pb.Message{
			MessageId: msg.MessageID,
			SenderId:  msg.SenderID,
			GroupId:   msg.GroupID,
			Text:      *msg.Text,
			CreatedAt: timestamppb.New(msg.CreatedAt),
		}
		if err := stream.Send(res); err != nil {
			return h.handleError(stream.Context(), "GetMessagesByGroup.Send", err)
		}
	}
	return nil
}
