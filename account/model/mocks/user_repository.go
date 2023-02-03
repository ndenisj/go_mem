package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/ndenisj/go_mem/account/model"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock type for model.UserRepository
type MockUserRepository struct {
	mock.Mock
}

// FindByID is mock of UserRepository FindByID
func (m *MockUserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	ret := m.Called(ctx, uid)

	var r0 *model.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// FindByID is mock of UserRepository FindByID
func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	ret := m.Called(ctx, email)

	var r0 *model.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// Create is mock of UserRepository Create
func (m *MockUserRepository) Create(ctx context.Context, u *model.User) error {
	ret := m.Called(ctx, u)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockUserRepository) Update(ctx context.Context, u *model.User) error {
	ret := m.Called(ctx, u)

	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockUserRepository) UpdateImage(ctx context.Context, uid uuid.UUID, imageURL string) (*model.User, error) {
	ret := m.Called(ctx, uid, imageURL)

	var r0 *model.User

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
