package entity

import "context"

// Location - courier location struct
type Location struct {
	UID       string
	CourierID string
	Lat       float32
	Lng       float32
}

// LocationRepository - interface describes location's repository responsibilities
type LocationRepository interface {
	InsertLocation(ctx context.Context, location Location) error
}
