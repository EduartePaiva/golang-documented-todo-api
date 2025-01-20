package session

import (
	"context"

	"github.com/golang-documented-todo-api/app/repository"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	CreateSession(ctx context.Context, arg repository.CreateSessionParams) error
}

type service struct {
	repository repository.Queries
}

// NewService is used to create a single instance of the service
func NewService(r repository.Queries) Service {
	return &service{
		repository: r,
	}
}

// InsertBook is a service layer that helps insert book in BookShop
func (s *service) CreateSession(ctx context.Context, arg repository.CreateSessionParams) error {
	return s.repository.CreateSession(ctx, arg)
}
