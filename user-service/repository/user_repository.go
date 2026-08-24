package repository

import (
	"context"

	"github.com/mohamed/microservices/user-service/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}


func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*entity.User, error) {

    var user entity.User

    if err := r.db.WithContext(ctx).
        First(&user, id).Error; err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (error) {

    if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
        return err
    }

    return nil
}