package unit

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserResponse struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Category string             `json:"Category" bson:"Category"`
	Qty      int                `json:"qty" bson:"qty"`
}
