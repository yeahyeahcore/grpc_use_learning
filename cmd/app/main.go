package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/config"
	grpc_controller "gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/interface/controller/grpc"
	"gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/interface/repository/mongodb"
	"gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/transport/grpc"

	"gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/usecase"
	"gitlab.doslab.ru/sell-and-buy/sb-delivery/pkg/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	// init configuration
	config, err := utils.ParseENV[config.Config]()
	if err != nil {
		log.Fatalln("config read is failed: ", err)
	}

	// init logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// init db connection
	mongoConnection, err := utils.GetMongoConnection(&config.Database)
	if err != nil {
		log.Fatalln("failed connection to db: ", err)
	}

	// init repositories
	locationRepository := mongodb.NewLocationRepo(mongoConnection)

	// init services
	locationService := usecase.NewLocationInteractor(locationRepository)

	// init gRPC
	grpcServer := grpc.New(logger).Register(grpc.Deps{
		LocationController: grpc_controller.NewLocationController(locationService),
	})

	go runGRPC(grpcServer, logger)

	//graceful shutdown
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Infoln("Shutting down server ...")
	grpcServer.Stop()
}

func runGRPC(grpcServer *grpc.Server, logger *logrus.Logger) {
	logger.Infoln("Starting gRPC Server")

	if err := grpcServer.Listen("localhost:8080"); err != nil {
		logger.Println("grpc Listen err: ", err)
	}
}
