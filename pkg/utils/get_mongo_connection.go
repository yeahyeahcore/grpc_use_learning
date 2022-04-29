package utils

import (
	"context"
	"fmt"

	"gitlab.doslab.ru/sell-and-buy/sb-delivery/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMongoConnection returns connection object to mongoDB
func GetMongoConnection(databaseConfig *config.DatabaseConfiguration) (*mongo.Client, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Auth,
	)

	return mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
}
