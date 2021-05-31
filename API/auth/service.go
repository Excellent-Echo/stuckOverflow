package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	// err = godotenv.Load()
	key = os.Getenv("SECRET_KEY")
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// if err != nil {
	// 	panic(err)
	// }

	claim := jwt.MapClaims{
		"user_id": userID,
	}

	// generate token using HS256 with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
