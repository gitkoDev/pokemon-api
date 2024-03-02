package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/gitkoDev/pokemon-db/models"
	"github.com/gitkoDev/pokemon-db/pkg/repository"
)

var salt = "ds46ghfgdbvcvbqqz555"

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

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
