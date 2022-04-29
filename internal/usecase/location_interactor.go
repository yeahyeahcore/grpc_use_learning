package usecase

import "gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/entity"

// LocationInteractor responsible for the business logic layer of location model
type LocationInteractor struct {
	locationRepo entity.LocationRepository
}

// NewLocationInteractor returns new instance of LocationInteractor
func NewLocationInteractor(locationRepository entity.LocationRepository) *LocationInteractor {
	return &LocationInteractor{
		locationRepo: locationRepository,
	}
}
