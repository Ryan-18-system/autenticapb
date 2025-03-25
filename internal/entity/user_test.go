package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("ryan", "123456", "ryanbrandao18@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.ID)
	assert.Equal(t, user.Username, "ryan")
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, user.Email, "ryanbrandao18@gmail.com")
	assert.True(t, user.ComparePassword("123456"))
}
func TestUser_ComparePassword(t *testing.T) {
	user, err := NewUser("ryan", "123456", "ryanbrandao18@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ComparePassword("123456"))
	assert.False(t, user.ComparePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
