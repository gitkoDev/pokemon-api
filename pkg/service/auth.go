package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/gitkoDev/pokemon-db/models"
	"github.com/gitkoDev/pokemon-db/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "ds46ghfgdbvcvbqqz555"
	signingKey = "58gg4409b48bqwtyhbv7"
	tokenTTL   = time.Hour * 12
)

type tokenClaims struct {
	jwt.RegisteredClaims
	TrainerId int `json:"trainer_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateTrainer(trainer models.Trainer) (int, error) {
	trainer.Password = s.generatePasswordHash(trainer.Password)
	return s.repo.CreateTrainer(trainer)
}

func (s *AuthService) GetTrainer(name string, password string) (models.Trainer, error) {
	password = s.generatePasswordHash(password)
	return s.repo.GetTrainer(name, password)
}

func (s *AuthService) GenerateToken(name string, password string) (string, error) {
	// Check if the trainer exists
	trainer, err := s.repo.GetTrainer(name, s.generatePasswordHash(password))
	if err != nil {
		return "", nil
	}

	// If trainer is found, create claims for jwt token
	claims := tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		trainer.Id,
	}

	// Create token and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	// Parse received token
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	// Token validation
	// 1. General errors checking
	if err != nil {
		return 0, err
		// 2. Checking for token signing method
	} else if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return 0, errors.New("wrong token signing method")
		// 3. Validating the presence of all claims
	} else if claims, ok := token.Claims.(*tokenClaims); !ok {
		return 0, errors.New("parsed tokens are not of type tokenClaims")
	} else {
		// 4. If passed validation, return id of trainer from DB
		return claims.TrainerId, nil
	}
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
