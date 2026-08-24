package handler

import (
	"context"

	"github.com/mohamed/microservices/user-service/dto"
	userpb "github.com/mohamed/microservices/user-service/proto"
	"github.com/mohamed/microservices/user-service/service"
)

type UserHandler struct {
	userService *service.UserService

    userpb.UnimplementedUserServiceServer
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUserById(ctx context.Context, req *userpb.UserIdRequest) (*userpb.User, error) {

    user, err := h.userService.GetUserById(
        ctx,
        int(req.Id),
    )

    if err != nil {
        return nil, err
    }

	return &userpb.User{
        Id:    int32(user.ID),
        Name:  user.Name,
        Email: user.Email,
    }, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.User, error) {
		
    user, err := h.userService.CreateUser(
        ctx,
        dto.CreateUserRequest{
            Name:  req.Name,
            Email: req.Email,
        },
    )

    if err != nil {
        return nil, err
    }

    return &userpb.User{
        Id:    int32(user.ID),
        Name:  user.Name,
        Email: user.Email,
    }, nil
}
