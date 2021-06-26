package repository

import (
	"post/domain"

	"github.com/jinzhu/gorm"
)

type PostRepository interface {
	GetPosts() ([]*domain.Post, error)
	GetPostsByUserId(userId int64) ([]*domain.Post, error)
	GetPostById(id uint64) (*domain.Post, error)
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

func (r *postRepository) GetPostsByUserId(userId int64) ([]*domain.Post, error) {
	var posts []*domain.Post
	if err := r.db.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetPostById(id uint64) (*domain.Post, error) {
	post := &domain.Post{}
	if err := r.db.Preload("PostDetails").Take(&post, id).Error; err != nil {
		return nil, err
	}
	return post, nil
}
