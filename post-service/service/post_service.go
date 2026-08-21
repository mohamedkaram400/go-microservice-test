package service

import (
	"context"

	"github.com/mohamed/microservices/post-service/dto"
	"github.com/mohamed/microservices/post-service/entity"
	pb "github.com/mohamed/microservices/post-service/proto"
	"github.com/mohamed/microservices/post-service/repository"
)

type PostService struct {
	postRepository *repository.PostRepository
	pb.UnimplementedPostServiceServer

} 

func NewPostService(postRepository *repository.PostRepository) *PostService {
	return &PostService{
		postRepository: postRepository,
	}
}

func (s *PostService) GetAllPosts(ctx context.Context, req *pb.Empty) (*pb.PostList, error) {

	return nil, nil
}

func (s *PostService) GetPostById(ctx context.Context, req *pb.PostIdRequest) (*pb.Post, error) {

	return nil, nil
}

func (s *PostService) CreatePost(ctx context.Context, req dto.CreatePostRequest) (*entity.Post, error) {
	
	post := &entity.Post{
		UserID: req.UserID,
		Title: req.Title,
		Description: req.Description,
	}

	err := s.postRepository.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) DeletePost(ctx context.Context, req *pb.PostIdRequest) (*pb.Post, error) {

	return nil, nil
}