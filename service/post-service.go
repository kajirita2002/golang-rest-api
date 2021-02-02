package service

import (
	"errors"
	"github/kaji2002/entity"
	"github/kaji2002/repository"
	"math/rand"
)

// PostService型のものは実装しなければならない
type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

// PostRepositoryの抽象化したrepoを搭載したNewPostServiceを作成
func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

// 実際の機能実装
func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
