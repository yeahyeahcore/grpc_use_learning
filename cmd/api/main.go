package api

import (
	"os"
	"os/signal"
	"syscall"

	"grpc_use_tutor/internal/transport/grpc"
	grpc_handler "grpc_use_tutor/internal/transport/grpc/handler"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	grpcHandler := grpc_handler.New(logger)
	grpcServer := grpc.New(logger).Register(grpc.Deps{
		LocationHandler: grpcHandler.LocationHandler,
	})

	go func(){
		logger.Infoln("Starting gRPC Server")
		
		if err := grpcServer.Listen("localhost:8080"); err != nil {
			logger.Println("grpc Listen err: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Infoln("Shutting down server")
	grpcServer.Stop()
}