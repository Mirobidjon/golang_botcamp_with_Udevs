package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Mirobidjon/todo-app"
	"github.com/Mirobidjon/todo-app/pkg/repository"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	salt       = "kasjdkja65646dasdsdjk"
	signingKey = "abc45SHHSG"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	UserID int
	jwt.StandardClaims
}

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(c repository.Autorization) *AuthService {
	return &AuthService{repo: c}
}

// CreateUser create new user
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken ...
func (s *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	// fmt.Println("token keldi ukam \n\tfirst\n\n", err)
	if err != nil {
		return "", err
	}
	
	fmt.Println("everythings good !")

	tk := &tokenClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	def, err := token.SignedString([]byte(signingKey))
	fmt.Printf("token: %s \nerr: %v \nuser.ID: %d\n\n", def, err, user.ID)
	return def, err
}


// ParseToken ...
func (s *AuthService) ParseToken(accesToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
