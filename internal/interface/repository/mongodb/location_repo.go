package mongodb

import (
	"context"

	"gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

// LocationRepo - repository of Location's collection
type LocationRepo struct {
	conn *mongo.Client
}

// NewLocationRepo returns new instance of Lication's repository
func NewLocationRepo(mongoConnection *mongo.Client) entity.LocationRepository {
	return &LocationRepo{
		conn: mongoConnection,
	}
}

// InsertLocation - insert new location
func (receiver *LocationRepo) InsertLocation(ctx context.Context, location entity.Location) error {
	return nil
}
