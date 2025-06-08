package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	grpc_controller "MuchUp/backend/internal/controllers/gprc/v2"
	rest_controller "MuchUp/backend/internal/controllers/http/rest"
	ws_controller "MuchUp/backend/internal/controllers/ws"
	"MuchUp/backend/internal/infrastructure/database"
	"MuchUp/backend/internal/infrastructure/database/repositories"

	group_service "MuchUp/backend/internal/application/service/group"
	message_service "MuchUp/backend/internal/application/service/message"
	user_service "MuchUp/backend/internal/application/service/user"

	"MuchUp/backend/pkg/logger"
	pb "MuchUp/backend/proto/gen/go/v2"

	"google.golang.org/grpc"
)

func main() {
	// 2. Setup Logger
	appLogger := logger.NewLogger() // Assuming a simple constructor

	// 3. Connect to Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		appLogger.Fatal("Failed to connect to database", err)
	}

	// 4. Run Migrations
	database.InitDB(db)
	appLogger.Info("Database migration completed")

	// 5. Initialize Repositories
	userRepo := repositories.NewUserRepository(db)
	messageRepo := repositories.NewMessageRepository(db)
	groupRepo := repositories.NewChatGroupRepository(db)

	// 6. Initialize Usecases (Services)
	groupUsecase := group_service.NewGroupUsecase(groupRepo, appLogger)
	userUsecase := user_service.NewUserUsecase(userRepo, groupUsecase)
	messageUsecase := message_service.NewMessageUsecase(messageRepo)

	// 7. Initialize Handlers
	restHandler := rest_controller.NewHandler(userUsecase, messageUsecase, appLogger)
	wsHandler := ws_controller.NewChatHandler(messageUsecase, userUsecase)
	grpcHandler := grpc_controller.NewGrpcHandler(userUsecase, messageUsecase, appLogger)

	// 8. Start Servers (gRPC and HTTP)
	// Start gRPC server in a separate goroutine
	go func() {
		grpcPort := os.Getenv("GRPC_PORT")
		if grpcPort == "" {
			grpcPort = "50051"
		}
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			appLogger.Fatal("Failed to listen for gRPC", err)
		}
		s := grpc.NewServer()
		pb.RegisterUserServiceServer(s, grpcHandler)
		pb.RegisterMessageServiceServer(s, grpcHandler)
		appLogger.Info("gRPC server listening at " + lis.Addr().String())
		if err := s.Serve(lis); err != nil {
			appLogger.Fatal("Failed to serve gRPC", err)
		}
	}()

	// Start HTTP server (blocking)
	router := restHandler.SetupRoutes()
	router.HandleFunc("/ws/chat", wsHandler.HandleWebSocket)

	httpPort := os.Getenv("SERVER_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	serverAddr := fmt.Sprintf(":%s", httpPort)
	appLogger.Info("HTTP server starting on " + serverAddr)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		appLogger.Fatal("Failed to start HTTP server", err)
	}
}
