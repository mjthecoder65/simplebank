package token

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mjthecoder65/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	jwtMaker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := faker.Username()
	duration := time.Minute

	token, err := jwtMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := jwtMaker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, claims)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := faker.Username()
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, claims)
}

func TestInvalidJWTToken(t *testing.T) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone,
		jwt.MapClaims{
			"username": faker.Username(),
			"exp":      time.Now().Add(time.Minute).Unix(),
		},
	)

	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	claims, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, claims)
}
