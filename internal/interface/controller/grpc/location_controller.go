package grpc

import (
	"gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/usecase"

	"github.com/yeahyeahcore/grpc_tutor/api"
)

// LocationController ...
type LocationController struct {
	interactor *usecase.LocationInteractor
	api.UnimplementedLocationServer
}

// NewLocationController returns new instance of LocationController
func NewLocationController(locationInteractor *usecase.LocationInteractor) *LocationController {
	return &LocationController{interactor: locationInteractor}
}
