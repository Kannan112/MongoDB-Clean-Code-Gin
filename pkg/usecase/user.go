package usecase

import (
	"context"
	"fmt"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"github.com/kannan112/go-gin-clean-arch/pkg/unit"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) Save(ctx context.Context, user unit.UserRequest) error {
	if err := c.userRepo.Save(ctx, user); err != nil {
		return fmt.Errorf("some thing failed :%v", err)
	}
	return nil
}

func (c *userUseCase) Read(ctx context.Context) ([]unit.UserResponse, error) {
	data, err := c.userRepo.Read(ctx)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (c *userUseCase) Update(ctx context.Context, productId int, data unit.UserRequest) error {
	err := c.userRepo.Update(ctx, productId, data)
	if err != nil {
		return fmt.Errorf("failed to update data: %v", err)
	}
	return nil
}

func (c *userUseCase) Delete(ctx context.Context, productId primitive.ObjectID) error {
	err := c.userRepo.Delete(ctx, productId)
	if err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}
