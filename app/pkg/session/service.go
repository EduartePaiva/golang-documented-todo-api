package session

import (
	"context"

	"github.com/golang-documented-todo-api/app/datasources/db"
	"github.com/golang-documented-todo-api/app/repository"
)

// Service is an interface from which our api module can access our repository of all our models
type SessionService interface {
	CreateSession(ctx context.Context, arg repository.CreateSessionParams) error
	SelectUserBySessionID(ctx context.Context, id string) (repository.SelectUserBySessionIDRow, error)
	UpdateSessionExpiresAt(ctx context.Context, arg repository.UpdateSessionExpiresAtParams) error
}

type sessionService struct {
	db db.Database
}

// NewService is used to create a single instance of the service
func NewService(db db.Database) SessionService {
	return &sessionService{db: db}
}

// InsertBook is a service layer that helps insert book in BookShop
func (s *sessionService) CreateSession(ctx context.Context, arg repository.CreateSessionParams) error {
	return s.db.CreateSession(ctx, arg)
}

func (s *sessionService) SelectUserBySessionID(ctx context.Context, id string) (repository.SelectUserBySessionIDRow, error) {
	return s.db.SelectUserBySessionID(ctx, id)
}

func (s *sessionService) UpdateSessionExpiresAt(ctx context.Context, arg repository.UpdateSessionExpiresAtParams) error {
	return s.db.UpdateSessionExpiresAt(ctx, arg)
}
