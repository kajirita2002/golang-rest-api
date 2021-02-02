package service

import (
	"errors"
	"math/rand"
	"github/kaji2002/entity"
	"github/kaji2002/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func NewPostService() PostService {
	return &service
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
	}
	if post.Title == "" {
		err := errors.New("The post is empty")
		return nil
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) Validate(post *entity.Post) ([]entity.Post, error) {
	if post ==
}