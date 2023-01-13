package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ndenisj/go_mem/account/model"
)

// UserService acts as a struct for injecting an implementation of
// UserRepository for use in service methods
type UserService struct {
	UserRepository model.UserRepository
}

// USConfig will hold repository that will eventually be injected
// into this service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factory function for initializing
// a UserService with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

// Get retrieves a user based on thier uuid
func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)

	return u, err
}