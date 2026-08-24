package repository

import (
	"context"

	"github.com/mohamed/microservices/post-service/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
    db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
    return &PostRepository{
        db: db,
    }
}

func (r *PostRepository) GetAllPosts(ctx context.Context) ([]*entity.Post, error) {
    var posts []*entity.Post
    if err := r.db.WithContext(ctx).Find(&posts).Error; err != nil {
        return nil, err
    }

    return posts, nil
}

func (r *PostRepository) GetPostById(ctx context.Context, postId int) (*entity.Post, error) {
    var post entity.Post

    if err := r.db.WithContext(ctx).Where("id = ?", postId).First(&post).Error; err != nil { 
        return nil, err
    }
    return &post, nil
}

func (r *PostRepository) Create(ctx context.Context, post *entity.Post) (error) {

    if err := r.db.WithContext(ctx).Create(post).Error; err != nil {
        return err
    }

    return nil
}

func (r *PostRepository) Delete(ctx context.Context, postId int) (error) {

    if err := r.db.WithContext(ctx).Where("id = ?", postId).Delete(&entity.Post{}).Error; err != nil { 
        return err
    }
    return nil
}
