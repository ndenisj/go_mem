package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/ndenisj/go_mem/account/model"
	"github.com/ndenisj/go_mem/account/model/apperrors"
)

// userService acts as a struct for injecting an implementation of
// UserRepository for use in service methods
type userService struct {
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
	return &userService{
		UserRepository: c.UserRepository,
	}
}

// Get retrieves a user based on thier uuid
func (s *userService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)

	return u, err
}

// Signup reaches our UserRepository to verify the email address is
// available and signs up the user if that is the case
func (s *userService) Signup(ctx context.Context, u *model.User) error {
	pw, err := hashPassword(u.Password)

	if err != nil {
		log.Printf("unable to signup user for email: %v\n", u.Email)
		return apperrors.NewInternal()
	}

	// now I realize why I originally used Signup(ctx, email, password)
	// then created a user. It's somewhat un-natural to mutate the user here
	u.Password = pw

	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}

	// If we get around to add events, we'd provide it here
	// err := s.EventsBroker.PublishUserUpdated(u, true)

	// if err != nil {
	// 	return nil, apperrors.NewInternal()
	// }

	return nil
}

// Signin reaches out to the UserRepository, checks if the user exists
// and then compares the supplied password with the provided password
// if a valid email/password is provided, you will hold all the available
// user fields
func (s *userService) Signin(ctx context.Context, u *model.User) error {
	uFetched, err := s.UserRepository.FindByEmail(ctx, u.Email)

	// Will return NotAuthorized to client to omit details of why
	if err != nil {
		return apperrors.NewAuthorization("Invalid email and password combination")
	}

	// verify password - we previously created this method
	match, err := comparePasswords(uFetched.Password, u.Password)

	if err != nil {
		return apperrors.NewInternal()
	}

	if !match {
		return apperrors.NewAuthorization("Invalid email and password combination")
	}

	*u = *uFetched

	return nil
}

func (s *userService) UpdateDetails(ctx context.Context, u *model.User) error {
	// Update user in UserRepository
	err := s.UserRepository.Update(ctx, u)
	if err != nil {
		return err
	}

	// Publish user updated
	// err = s.EventsBroker.PublishUserUpdated(u, false)
	// if err != nil {
	// 	return apperrors.NewInternal()
	// }

	return nil
}
