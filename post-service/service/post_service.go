package service

import (
	"context"

	pb "github.com/mohamed/microservices/post-service/proto"
)

type PostService struct {
	pb.UnimplementedPostServiceServer

	posts []*pb.Post
}

func NewPostService() *PostService {
	return &PostService{
		posts: []*pb.Post{
			{
				Id: 1,
				Title: "title 1",
				Description: "description 1",
			},
			{
				Id: 1,
				Title: "title 2",
				Description: "description 2",
			},
		},
	}
}

func (s *PostService) GetAllUsers(
	ctx context.Context,
	req *pb.Empty,
) (*pb.PostList, error) {

	return &pb.PostList{
		Posts: s.posts,
	}, nil
}

func (s *PostService) GetUserById(
	ctx context.Context,
	req *pb.PostIdRequest,
) (*pb.Post, error) {

	for _, post := range s.posts {
		if post.Id == req.Id {
			return post, nil
		}
	}

	return nil, nil
}
