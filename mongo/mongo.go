package mongo

import (
	"context"
	"log"

	"example.com/demo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	client *mongo.Client
	config *demo.MongoConfig
	log    *log.Logger
}

func NewDB(config *demo.MongoConfig, log *log.Logger) *db {
	return &db{
		config: config,
		log:    log,
	}
}

func (db *db) Open() error {
	mongoUri := db.config.ConnectionUrl
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		db.log.Println("Failed to connect to MongoDB", err)
		return err
	}

	db.client = client
	return nil
}
