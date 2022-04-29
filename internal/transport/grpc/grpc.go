package grpc

import (
	"net"

	grpc_controller "gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/interface/controller/grpc"

	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/grpc_tutor/api"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

// Deps ...
type Deps struct {
	LocationController *grpc_controller.LocationController
}

// Server ...
type Server struct {
	Logger *logrus.Entry
	grpc   *grpc.Server
}

// New returns new instance of gRPC server
func New(logger *logrus.Logger) *Server {
	grpcLogger := logrus.NewEntry(logger)

	grpcStreamInterceptor := grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_logrus.StreamServerInterceptor(grpcLogger),
		grpc_recovery.StreamServerInterceptor(),
	))
	grpcUnaryInterceptor := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(grpcLogger),
		grpc_recovery.UnaryServerInterceptor(),
	))

	return &Server{
		grpc:   grpc.NewServer(grpcStreamInterceptor, grpcUnaryInterceptor),
		Logger: grpcLogger,
	}
}

// Listen requests
func (receiver *Server) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	if err := receiver.grpc.Serve(listener); err != nil {
		return err
	}

	return nil
}

// Register gRPC controllers
func (receiver *Server) Register(deps Deps) *Server {
	api.RegisterLocationServer(receiver.grpc, deps.LocationController)

	return receiver
}

// Stop gRPC server
func (receiver *Server) Stop() {
	receiver.grpc.GracefulStop()
}
