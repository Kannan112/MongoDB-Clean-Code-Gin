package repository

import (
	"context"
	"fmt"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"github.com/kannan112/go-gin-clean-arch/pkg/unit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDatabase struct {
	DB *mongo.Client
}

func NewUserRepository(DB *mongo.Client) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (db *userDatabase) Save(ctx context.Context, user unit.UserRequest) error {
	collections := db.DB.Database("mongo_data").Collection("users")
	res, err := collections.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("error insertion user:%v", err)
	}
	fmt.Printf("res: %v\n", res)
	return err
}

func (db *userDatabase) Read(ctx context.Context) ([]unit.UserResponse, error) {
	var value []unit.UserResponse
	collection := db.DB.Database("mongo_data").Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return value, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user unit.UserResponse
		err := cursor.Decode(&user)
		if err != nil {
			return value, err
		}
		value = append(value, user)
	}
	if err := cursor.Err(); err != nil {
		return value, err
	}

	return value, nil
}

func (db *userDatabase) Update(ctx context.Context, productId int, data unit.UserRequest) error {
	collection := db.DB.Database("mongo_data").Collection("users")
	_, err := collection.UpdateOne(ctx, productId, data)
	return err
}
func (db *userDatabase) Delete(ctx context.Context, productId primitive.ObjectID) error {
	collection := db.DB.Database("mongo_data").Collection("users")
	_, err := collection.DeleteOne(ctx, productId)
	return err
}
