package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	config "github.com/kannan112/go-gin-clean-arch/pkg/config"
)

func ConnectDatabase(cfg config.Config) (*mongo.Client, error) {
	ctx := context.TODO()

	// Create a new MongoDB client
	connectionString := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server to check the connection
	// err = client.Ping(context.Background(), nil)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to ping MongoDB %v:", err)
	// }
	// database := client.Database("mongo_demo")
	// fmt.Println("database", database)

	// list, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to list database name :%v", err)
	// }
	// dbExists := false
	// for _, dbName := range list {
	// 	if dbName == "mongo_demo" {
	// 		dbExists = true
	// 		break
	// 	}
	// }
	// if !dbExists {
	// 	//create collection
	// 	if err := database.CreateCollection(ctx, "users"); err != nil {
	// 		return nil, fmt.Errorf("failed to create collection :%v", err)
	// 	}
	// } else {
	// 	log.Fatal("connected to database :", "mongo_db")
	// }
	return client, nil
}
