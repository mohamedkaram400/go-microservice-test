package dto


type CreatePostRequest struct {
	UserId int32	`json:"user_id"`
    Title string	`json:"title"`
    Description string	`json:"description"`
}

type UserResponse struct {
    ID    int32  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type PostResponse struct {
    ID          int32       `json:"id"`
    Title       string       `json:"title"`
    Description string       `json:"description"`
    User        UserResponse `json:"user"`
}