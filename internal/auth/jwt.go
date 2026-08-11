package auth

import (
	"errors"
	"time"

	"github.com/aromalcode-prog/cab-share-backend/config"
	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("Invalid token")

type JWTManager struct {
	secret     []byte
	expiration time.Duration
}

func NewJWTManager(cfg *config.Config) *JWTManager {
	return &JWTManager{
		secret:     []byte(cfg.JWTSecret),
		expiration: time.Duration(cfg.JWTExpirationHours) * time.Hour,
	}
}
func (j *JWTManager) GenerateJWT(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  &jwt.NumericDate{time.Now()},
			ExpiresAt: &jwt.NumericDate{time.Now().Add(j.expiration)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(j.secret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWTManager) VerifyJWT(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.secret, nil
		},
	)
	if err != nil {
		return nil, ErrInvalidToken
	}
	claim, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidToken
	}
	return claim, nil
}
