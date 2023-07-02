package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/unit"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Save(ctx context.Context, user unit.UserRequest) error                  // CREATE
	Read(ctx context.Context) ([]unit.UserResponse, error)                  // READ
	Update(ctx context.Context, productId int, data unit.UserRequest) error // UPDATE
	Delete(ctx context.Context, productId primitive.ObjectID) error         // DELETE
}
