package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// UserService defines methods the handler layer expects any
// service it interact with to implement
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
}

// TokenService defines methods the handler layer expect to interact with
// in regards to producing jwt as string
type TokenService interface {
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
}

// UserRepository defines methods the service layer expects any
// repository it interact with to implement
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
	Create(ctx context.Context, u *User) error
}

// TokenRepository defines methods that it expects a repository it
// interact with to implement
type TokenRepository interface {
	SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
}
