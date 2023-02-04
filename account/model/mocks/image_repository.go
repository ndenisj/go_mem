package mocks

import (
	"context"
	"mime/multipart"

	"github.com/stretchr/testify/mock"
)

// MockImageRepository is a mock type of model.ImageRepository
type MockImageRepository struct {
	mock.Mock
}

// UpdateProfile is mock representation of ImageRepository UpdateProfile
func (m *MockImageRepository) UpdateProfile(ctx context.Context, objName string, imageFile multipart.File) (string, error) {
	// args that will be passed to 'Return' in the test, when function
	// is called with uid. Hence the name 'ret'
	ret := m.Called(ctx, objName, imageFile)

	// first value passed to return
	var r0 string
	if ret.Get(0) != nil {
		// we can return this if we know we wont be passing function to Return
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockImageRepository) DeleteProfile(ctx context.Context, objName string) error {
	ret := m.Called(ctx, objName)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
