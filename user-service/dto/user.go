package dto


type CreateUserRequest struct {
	UserId int32	`json:"user_id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserResponse struct {
    ID    int32  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
