package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/unit"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (c *UserHandler) Save(ctx *gin.Context) {
	var NewUser unit.UserRequest
	if err := ctx.Bind(&NewUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	err := c.userUseCase.Save(ctx, NewUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, unit.Response{
			Status:  false,
			Message: "failed to save user data",
			Errors:  err,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, unit.Response{
		Status:  true,
		Message: "successfully created new user",
		Errors:  nil,
		Data:    nil,
	})
}

func (c *UserHandler) Read(ctx *gin.Context) {
	data, err := c.userUseCase.Read(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, unit.Response{
			Status:  false,
			Message: "failed to read the data",
			Errors:  err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, unit.Response{
		Status:  true,
		Message: "Data Found",
		Errors:  nil,
		Data:    data,
	})
}

func (c *UserHandler) Update(ctx *gin.Context) {
	var UpdateData unit.UserRequest
	productStr := ctx.Param("productId")
	producesID, err := strconv.Atoi(productStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, unit.Response{
			Status:  false,
			Message: "failed to get the id",
			Errors:  err.Error(),
			Data:    nil,
		})
		return
	}
	err = ctx.Bind(&UpdateData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, unit.Response{
			Status:  false,
			Message: "failed to bind the data",
			Errors:  err.Error(),
			Data:    nil,
		})
		return
	}
	err = c.userUseCase.Update(ctx, producesID, UpdateData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, unit.Response{
			Status:  false,
			Message: "failed to update",
			Errors:  err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, unit.Response{
		Status:  true,
		Message: "updated",
		Errors:  nil,
		Data:    nil,
	})
}

func (c *UserHandler) Delete(ctx *gin.Context) {
	productStr := ctx.Param("productId")
	objectID, err := primitive.ObjectIDFromHex(productStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, unit.Response{
			Status:  false,
			Message: "Failed to get the ID",
			Errors:  err.Error(),
			Data:    nil,
		})
		return
	}

	err = c.userUseCase.Delete(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, unit.Response{
			Status:  false,
			Message: "Failed to delete the data",
			Errors:  err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, unit.Response{
		Status:  true,
		Message: "Data deleted successfully",
		Errors:  nil,
		Data:    nil,
	})
}
