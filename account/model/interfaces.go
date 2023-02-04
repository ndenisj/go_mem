package model

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

/**
SERVICES get request from the handler, it it needs a datasource, it sends it to the Repositories
*/

// UserService defines methods the handler layer expects any
// service it interact with to implement
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
	Signin(ctx context.Context, u *User) error
	UpdateDetails(ctx context.Context, u *User) error
	SetProfileImage(ctx context.Context, uid uuid.UUID, imageFileHeader *multipart.FileHeader) (*User, error)
	ClearProfileImage(ctx context.Context, uid uuid.UUID) error
}

// TokenService defines methods the handler layer expect to interact with
// in regards to producing jwt as string
type TokenService interface {
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
	Signout(ctx context.Context, uid uuid.UUID) error
	ValidateIDToken(tokenString string) (*User, error)
	ValidateRefreshToken(refreshTokenString string) (*RefreshToken, error)
}

/**
REPOSITORIES get request from the services and communicate with datasources
*/
// UserRepository defines methods the service layer expects any
// repository it interact with to implement
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	UpdateImage(ctx context.Context, uid uuid.UUID, imageURL string) (*User, error)
}

// TokenRepository defines methods that it expects a repository it
// interact with to implement
type TokenRepository interface {
	SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
	DeleteUserRefreshTokens(ctx context.Context, userID string) error
}

// ImageRepository defines methods it expects a repository it
// interact with to implement
type ImageRepository interface {
	UpdateProfile(ctx context.Context, objName string, imageFile multipart.File) (string, error)
	DeleteProfile(ctx context.Context, objName string) error
}
