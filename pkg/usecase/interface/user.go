package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/unit"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCase interface {
	Save(ctx context.Context, user unit.UserRequest) error
	Read(ctx context.Context) ([]unit.UserResponse, error)
	Update(ctx context.Context, productId int, data unit.UserRequest) error
	Delete(ctx context.Context, productId primitive.ObjectID) error
}
