package server

import (
	"MuchUp/backend/config"
	"MuchUp/backend/pkg/logger"
	"fmt"
	"net"
	"net/http"
	"os"
	grpc_controller "MuchUp/backend/internal/controllers/gprc/v2"
	pb "MuchUp/backend/proto/gen/go/v2"

	"google.golang.org/grpc"
)

func StartGRPCServer(cfg *config.Config, appLogger logger.Logger, grpcHandler  *grpc_controller.GrpcHandler) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
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
}

func StartRestServer(cfg *config.Config,appLogger logger.Logger, router *http.ServeMux) {
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