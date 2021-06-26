package repository

import (
	"post/domain"

	"github.com/jinzhu/gorm"
)

type PostRepository interface {
	GetPosts() ([]*domain.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) GetPosts() ([]*domain.Post, error) {
	var posts []*domain.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
