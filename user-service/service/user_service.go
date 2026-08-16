package service

import (
	"context"

	pb "github.com/mohamed/microservices/user-service/proto"
)

type UserService struct {
	pb.UnimplementedUserServiceServer

	users []*pb.User
}

func NewUserService() *UserService {
	return &UserService{
		users: []*pb.User{
			{
				Id: 1,
				Name: "Ali",
				Email: "ali@gmail.com",
			},
			{
				Id: 1,
				Name: "Ahmed",
				Email: "ahmed@gmail.com",
			},
		},
	}
}

func (s *UserService) GetAllUsers(
	ctx context.Context,
	req *pb.Empty,
) (*pb.UserList, error) {

	return &pb.UserList{
		Users: s.users,
	}, nil
}