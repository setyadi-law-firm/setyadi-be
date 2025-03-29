package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/setyadi-law-firm/setyadi-be/app/models"
)

type Util struct {
	config *models.Config
}

func NewUtil(config *models.Config) *Util {
	return &Util{config}
}

func (u *Util) GenerateTokenPair(user *User) (string, string, error) {
	now := time.Now()
	jwtExpiryInDays := u.config.JwtExpiryInDays

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "setyadi-be",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(jwtExpiryInDays) * 24 * time.Hour)),
			Subject:   user.Email,
		},
	})

	signedJwtToken, err := token.SignedString([]byte(u.config.JwtSecretKey))
	if err != nil {
		return "", "", fmt.Errorf("unable to sign access token: %w", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "setyadi-be",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(u.config.JwtRefreshExpiryInDays) * 24 * time.Hour)),
			Subject:   user.Email,
		},
	})

	signedRefreshToken, err := refreshToken.SignedString([]byte(u.config.JwtSecretKey))
	if err != nil {
		return "", "", fmt.Errorf("unable to sign refresh token: %w", err)
	}

	return signedJwtToken, signedRefreshToken, nil
}

func (u *Util) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (u *Util) ComparePassword(hashedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}

func (u *Util) ExtractJwtToken(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	authSplit := strings.Split(authorization, " ")
	if len(authSplit) != 2 || authSplit[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header")
	}
	return authSplit[1], nil
}

func (u *Util) ToJwtToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Method)
		}
		return []byte(u.config.JwtSecretKey), nil
	})
}
