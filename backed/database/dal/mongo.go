package dal

import (
	"context"
	"fmt"
	"time"

	"Backed/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database

func InitMongo(cfg config.DBConfig) error {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		cfg.MongoUser, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return err
	}
	Mongo = client.Database(cfg.MongoDBName)
	return nil
}
