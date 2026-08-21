package dto


type CreatePostRequest struct {
	UserID int32	`json:"user_id"`
    Title string	`json:"title"`
    Description string	`json:"description"`
}