package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))

type JWTService struct{}

func NewJWTService() *JWTService {
    return &JWTService{}
}

type Claims struct {
    UserID string `json:"user_id"`
    RoleID string `json:"role_id"`
    jwt.RegisteredClaims
}

func (j *JWTService) GenerateAccessToken(userID string, roleID string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        RoleID: roleID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func (j *JWTService) VerifyAccessToken(tokenStr string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    return claims, nil
}

func (j *JWTService) GenerateRefreshToken(userID string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

func (j *JWTService) VerifyRefreshToken(tokenStr string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return refreshSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    return claims, nil
}
