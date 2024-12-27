package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomPassword()
	passwordHash, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, passwordHash)
	require.NotEqual(t, password, passwordHash)

	require.True(t, IsValidPassword(password, passwordHash))

	wrongPassword := RandomPassword()
	require.NotEqual(t, password, wrongPassword)
	require.False(t, IsValidPassword(wrongPassword, passwordHash))
}
