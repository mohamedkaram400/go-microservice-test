package service

import (
	"context"
	"log"

	"github.com/mohamed/microservices/post-service/dto"
	"github.com/mohamed/microservices/post-service/entity"
	pb "github.com/mohamed/microservices/post-service/proto"
	"github.com/mohamed/microservices/post-service/repository"
	userpb "github.com/mohamed/microservices/user-service/proto"
)

type PostService struct {
	postRepository *repository.PostRepository
	userClient     userpb.UserServiceClient
} 

func NewPostService(postRepository *repository.PostRepository, userClient userpb.UserServiceClient) *PostService {
	return &PostService{
		postRepository: postRepository,
		userClient: userClient,
	}
}

func (s *PostService) GetAllPosts(ctx context.Context, req *pb.Empty) ([]*entity.Post, error) {
	
	return s.postRepository.GetAllPosts(ctx)
}

func (s *PostService) GetPostById(ctx context.Context, postId int) (*dto.PostResponse, error) {
	post, err := s.postRepository.GetPostById(ctx, postId)
	if err != nil {
		return nil, err
	}

	user, err := s.userClient.GetUserById(
        ctx,
        &userpb.UserIdRequest{
            Id: post.UserID,
        },
    )

	if err != nil {
		return nil, err
	}

	log.Printf(
		"USER FROM GRPC: ID=%d NAME=%s EMAIL=%s",
		user.Id,
		user.Name,
		user.Email,
	)
	return &dto.PostResponse{
        ID:          post.ID,
        Title:       post.Title,
        Description: post.Description,
        User: dto.UserResponse{
            ID:    user.Id,
            Name:  user.Name,
            Email: user.Email,
        },
    }, nil
}

func (s *PostService) CreatePost(ctx context.Context, req dto.CreatePostRequest) (*entity.Post, error) {
	_, err := s.userClient.GetUserById(
        ctx,
        &userpb.UserIdRequest{
            Id: req.UserId,
        },
    )

	if err != nil {
		return nil, err
	}

	post := &entity.Post{
		UserID: req.UserId,
		Title: req.Title,
		Description: req.Description,
	}

	if err := s.postRepository.Create(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) DeletePost(ctx context.Context, postId int) (error) {

	return s.postRepository.Delete(ctx, postId)
}

