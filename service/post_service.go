package service

import (
	"math/rand"
	"errors"
	"go-rest-api/entity"
	"go-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}


var (
	repo repository.PostRepository
)


func NewPostService(r repository.PostRepository) PostService {
	repo = r
	return &service{}
}

func (s *service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("post title is empty")
		return err
	}
	return nil
}

func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}