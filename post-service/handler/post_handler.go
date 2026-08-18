package handler

import (
	"encoding/json"
	"net/http"

	userpb "github.com/mohamed/microservices/user-service/proto"
)

type Post struct {
    ID   int
    UserID int
    Title string
    Description string
}


type PostHandler struct {
	userClient userpb.UserServiceClient
}

func NewPostHandler(userClient userpb.UserServiceClient) *PostHandler {
	return &PostHandler{
		userClient: userClient,
	}
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	// Temporary hardcoded post.
	// Later this will come from your Post database.
	post := Post{
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