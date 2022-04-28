package handler

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/grpc_tutor/api"
)

type Location interface {
	GetLocation(ctx context.Context, request *api.LocationRequest) (*api.LocationResponse, error)
}

type LocationHandler struct {
	api.UnimplementedLocationServer

	logger *logrus.Logger
}

func (receiver *LocationHandler) New(logger *logrus.Logger) *LocationHandler {
	return &LocationHandler{
		logger: logger,
	}
}

func (receiver *LocationHandler) GetLocation(ctx context.Context, request *api.LocationRequest) (*api.LocationResponse, error) {
	receiver.logger.Infoln("Handle GetLocation", "couriedID", request.GetCourierId())

	return &api.LocationResponse{
		Lat: 4.5000,
		Lng: 2.3333,
	}, nil
}