package usecase

import (
	"post/domain"
	"post/interface/repository"
)

type PostUsecase interface {
	GetPosts() ([]*domain.Post, error)
	GetPostsByUserId(userId int64) ([]*domain.Post, error)
}

type postUsecase struct {
	postRepository repository.PostRepository
}

func NewPostUsecase(ur repository.PostRepository) *postUsecase {
	return &postUsecase{
		postRepository: ur,
	}
}

func (u *postUsecase) GetPosts() ([]*domain.Post, error) {
	return u.postRepository.GetPosts()
}

func (u *postUsecase) GetPostsByUserId(userId int64) ([]*domain.Post, error) {
	return u.postRepository.GetPostsByUserId(userId)
}
