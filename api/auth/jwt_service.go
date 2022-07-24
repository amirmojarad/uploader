package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(username string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetMapClaims(token *jwt.Token) jwt.MapClaims
}

type authCustomClaims struct {
	Username string `json:"username"`
	jwt.Claims
}

type jwtService struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	secret := os.Getenv("SECRET_KEY")
	return &jwtService{
		secretKey: secret,
		issure:    "interface-service",
	}
}

func (service *jwtService) GenerateToken(username string) string {
	claims := &authCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if t, err := token.SignedString([]byte(service.secretKey)); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (service *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %+v", t.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func (service *jwtService) GetMapClaims(token *jwt.Token) jwt.MapClaims {
	return token.Claims.(jwt.MapClaims)
}
