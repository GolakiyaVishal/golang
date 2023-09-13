package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct {
	// GetById (int)
	// Save    (*User)
}

func (m *MockUserRepository) GetById(id int) (*User, error) {
	return &User{name: "Vishal"}, nil
}

func (m *MockUserRepository) Save(user *User) error {
	return nil
}

func TestUserService_GetUserById(t *testing.T) {
	mockRepo := MockUserRepository{}

	userService := NewUserService(mockRepo)

	user, err := userService.GetById(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}
