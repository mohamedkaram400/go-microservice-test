package entity

type Post struct {
    ID   int32		`json:"id"`
    UserID int32	`json:"user_id"`
    Title string	`json:"title"`
    Description string	`json:"description"`
}
