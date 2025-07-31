package main

import (
	"MuchUp/backend/config"
	group_service "MuchUp/backend/internal/application/service/group"
	message_service "MuchUp/backend/internal/application/service/message"
	user_service "MuchUp/backend/internal/application/service/user"
	grpc_controller "MuchUp/backend/internal/controllers/grpc/v2"
	rest_controller "MuchUp/backend/internal/controllers/http/rest"
	ws_controller "MuchUp/backend/internal/controllers/ws"
	"MuchUp/backend/internal/infrastructure/database"
	"MuchUp/backend/internal/infrastructure/database/repositories"
	"MuchUp/backend/pkg/logger"
	"MuchUp/backend/pkg/middleware"

	"MuchUp/backend/internal/infrastructure/auth"
	"MuchUp/backend/internal/infrastructure/server"
	"net/http"

	"github.com/gorilla/mux"
)

// @Title MuchUp API
// @version 1.0
// @description This is the MuchUp server
// @host localhost:8080
// @Basepath /api/v1
func main() {
	config := config.LoadConfig()
	appLogger := logger.NewLogger()

	db, err := database.Connect(config)
	if err != nil {
		appLogger.Fatal("Failed to connect to database", err)
	}

	database.InitDB(db)
	appLogger.Info("Database migration completed")
	JWTValidator := auth.NewJWTValidator(config.SecretKey, "Much-Up", "users")
	userRepo := repositories.NewUserRepository(db)
	messageRepo := repositories.NewMessageRepository(db)
	groupRepo := repositories.NewChatGroupRepository(db)
	groupUsecase := group_service.NewGroupUsecase(groupRepo, appLogger)
	userUsecase := user_service.NewUserUsecase(userRepo, groupUsecase)
	messageUsecase := message_service.NewMessageUsecase(messageRepo, userRepo)
	chatHandler := ws_controller.NewChatHandler(messageUsecase, userUsecase)
	wsHandler := http.HandlerFunc(chatHandler.HandleWebSocket)
	RestHandler := rest_controller.NewHandler(userUsecase, messageUsecase, appLogger)

	grpcHandler := grpc_controller.NewGrpcHandler(userUsecase, messageUsecase, appLogger)

	wsRouter := mux.NewRouter()
	authenticatedWsHandler := middleware.JWTMiddleware(wsHandler, JWTValidator)
	wsRouter.Handle("/ws/chat", authenticatedWsHandler)

	go server.StartGRPCServer(config, appLogger, grpcHandler)
	go server.StartHTTPServer(config, appLogger, RestHandler.SetupRoutes(JWTValidator))

}
