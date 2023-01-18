package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/ndenisj/go_mem/account/model"
	"github.com/ndenisj/go_mem/account/model/apperrors"
	"github.com/ndenisj/go_mem/account/model/mocks"
	"github.com/ndenisj/go_mem/account/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run(
		"Success",
		func(t *testing.T) {
			uid, _ := uuid.NewRandom()

			mockUserResp := &model.User{
				UID:   uid,
				Email: util.RandomEmail(),
				Name:  util.RandomFullname(),
			}

			mockUserRepository := new(mocks.MockUserRepository)
			us := NewUserService(&USConfig{
				UserRepository: mockUserRepository,
			})
			mockUserRepository.On("FindByID", mock.Anything, uid).Return(mockUserResp, nil)

			ctx := context.TODO()
			u, err := us.Get(ctx, uid)

			assert.NoError(t, err)
			assert.Equal(t, u, mockUserResp)
			mockUserRepository.AssertExpectations(t)
		},
	)

	t.Run(
		"Error",
		func(t *testing.T) {
			uid, _ := uuid.NewRandom()

			mockUserRepository := new(mocks.MockUserRepository)
			us := NewUserService(&USConfig{
				UserRepository: mockUserRepository,
			})

			mockUserRepository.On("FindByID", mock.Anything, uid).Return(nil, fmt.Errorf("some error down the call chain"))

			ctx := context.TODO()
			u, err := us.Get(ctx, uid)

			assert.Nil(t, u)
			assert.Error(t, err)
			mockUserRepository.AssertExpectations(t)
		},
	)
}

func TestSignup(t *testing.T) {
	t.Run(
		"Success",
		func(t *testing.T) {
			uid, _ := uuid.NewRandom()

			mockUser := &model.User{
				Email:    "john@doe.com",
				Password: "12err434ssss",
			}

			mockUserRepository := new(mocks.MockUserRepository)
			us := NewUserService(&USConfig{
				UserRepository: mockUserRepository,
			})

			// we can use Run method modify the user when the create method is called
			// we can then chain on a Return method to return no error
			mockUserRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).Run(
				func(args mock.Arguments) {
					userArg := args.Get(1).(*model.User) // arg 0 is context, arg 1 is *User
					userArg.UID = uid
				},
			).Return(nil)

			ctx := context.TODO()
			err := us.Signup(ctx, mockUser)

			assert.NoError(t, err)

			assert.Equal(t, uid, mockUser.UID)

			mockUserRepository.AssertExpectations(t)
		},
	)

	t.Run(
		"Error",
		func(t *testing.T) {
			mockUser := &model.User{
				Email:    "john@doe.com",
				Password: "12err434ssss",
			}

			mockUserRepository := new(mocks.MockUserRepository)
			us := NewUserService(&USConfig{
				UserRepository: mockUserRepository,
			})

			mockErr := apperrors.NewConflict("email", mockUser.Email)

			mockUserRepository.On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).Return(mockErr)

			ctx := context.TODO()
			err := us.Signup(ctx, mockUser)

			assert.EqualError(t, err, mockErr.Error())
			mockUserRepository.AssertExpectations(t)
		},
	)
}