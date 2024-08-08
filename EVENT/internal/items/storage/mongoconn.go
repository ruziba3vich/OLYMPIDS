package storage

import (
	"context"
	"fmt"

	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client          *mongo.Client
	StatsCollection *mongo.Collection
}

func ConnectMongoDB(cfg *config.Config, ctx context.Context) (*DB, error) {
	clientOptions := options.Client().ApplyURI(cfg.Mongosh.MongoUri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %s", err.Error())
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %s", err.Error())
	}

	return &DB{
		Client:          client,
		StatsCollection: client.Database(cfg.Mongosh.MongoDb).Collection(cfg.Mongosh.Collection),
	}, nil
}

// DisconnectDB to disconnect the db
func (db *DB) DisconnectDB(ctx context.Context) error {
	if err := db.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %s", err.Error())
	}
	return nil
}
