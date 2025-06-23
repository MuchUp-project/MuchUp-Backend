package main

import (
	group_service "MuchUp/backend/internal/application/service/group"
	message_service "MuchUp/backend/internal/application/service/message"
	user_service "MuchUp/backend/internal/application/service/user"
	grpc_controller "MuchUp/backend/internal/controllers/gprc/v2"
	ws_controller "MuchUp/backend/internal/controllers/ws"
	"MuchUp/backend/internal/infrastructure/database"
	"MuchUp/backend/internal/infrastructure/database/repositories"
	"MuchUp/backend/config"
	"MuchUp/backend/pkg/logger"
	"MuchUp/backend/pkg/middleware"

	"MuchUp/backend/internal/infrastructure/server"
	"MuchUp/backend/internal/infrastructure/auth"
	"net/http"
)

func main() {
	config := config.LoadConfig()
	appLogger := logger.NewLogger()
	
	db,err := database.Connect(config)
	if err != nil {
		appLogger.Fatal("Failed to connect to database", err)
	}
	database.InitDB(db)
	appLogger.Info("Database migration completed")
	JWTValidator :=  auth.NewJWTValidator(config.SecretKey,"Much-Up","users")
	userRepo := repositories.NewUserRepository(db)
	messageRepo := repositories.NewMessageRepository(db)
	groupRepo := repositories.NewChatGroupRepository(db)
	groupUsecase := group_service.NewGroupUsecase(groupRepo, appLogger)
	userUsecase := user_service.NewUserUsecase(userRepo, groupUsecase)
	messageUsecase := message_service.NewMessageUsecase(messageRepo)
	chatHandler := ws_controller.NewChatHandler(messageUsecase,userUsecase)
	wsHandler := http.HandlerFunc(chatHandler.HandleWebSocket)
	

	
	grpcHandler := grpc_controller.NewGrpcHandler(userUsecase, messageUsecase, appLogger)

	mux := http.NewServeMux()
	authenticatedWsHandler := middleware.JWTMiddleware(wsHandler,JWTValidator)
	mux.Handle("/ws/chat",authenticatedWsHandler)
	
	go server.StartGRPCServer(config,appLogger,grpcHandler)
	go server.StartRestServer(config,appLogger,mux)
	
}