package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (jwt.Claims, error)
}
