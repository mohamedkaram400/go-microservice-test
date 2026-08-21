package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mohamed/microservices/post-service/dto"
	"github.com/mohamed/microservices/post-service/entity"
	"github.com/mohamed/microservices/post-service/service"
	userpb "github.com/mohamed/microservices/user-service/proto"
)

type PostHandler struct {
	userClient userpb.UserServiceClient
	postService *service.PostService
}

func NewPostHandler(userClient userpb.UserServiceClient, postService *service.PostService) *PostHandler {
	return &PostHandler{
		userClient: userClient,
		postService: postService,
	}
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {

}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	// Temporary hardcoded post.
	// Later this will come from your Post database.
	post := entity.Post{
		ID:          1,
		UserID:      1,
		Title:       "My first post",
		Description: "Hello from Post Service",
	}

	// Call User Service through gRPC.
	user, err := h.userClient.GetUserById(
		r.Context(),
		&userpb.UserIdRequest{
			Id: int32(post.UserID),
		},
	)

	if err != nil {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":          post.ID,
		"title":       post.Title,
		"description": post.Description,
		"user": map[string]interface{}{
			"id":    user.Id,
			"name":  user.Name,
			"email": user.Email,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
		
	var req dto.CreatePostRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	post, err := h.postService.CreatePost(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

}