package repository

import (
	"context"
	"database/sql"

	"github.com/mohamed/microservices/post-service/entity"
)

type PostRepository struct {
    db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
    return &PostRepository{
        db: db,
    }
}

func (r *PostRepository) GetAllPosts(ctx context.Context) (error) {
    
    return nil
}

func (r *PostRepository) GetPostById(ctx context.Context, postId int) (error) {

    return nil
}

func (r *PostRepository) Create(ctx context.Context, post *entity.Post) (error) {
    query := `
        INSERT INTO posts (
            user_id,
            title, 
            description,
            created_at,
            updated_at
        )
        VALUES (?, ?, ?, NOW(), NOW())
    `

    _, err := r.db.ExecContext(
        ctx, 
        query, 
        post.UserID,
        post.Title,
        post.Description,
    )

    if err != nil {
        return err
    }

    return nil
}

func (r *PostRepository) Delete(ctx context.Context, postId int) (error) {

    return nil
}
