package service

import (
	"context"
	"errors"
	"time"

	"casino.sandboxforge.tech/casino-api/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type User = domain.User

var (
	ErrLoginTaken = errors.New("login taken")
	ErrEmailTaken = errors.New("email taken")
)

type UserRepo interface {
	ByLogin(ctx context.Context, login string) (*User, error)
	ByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, login, email, pwdHash string) (int64, error)
}

type Config interface {
	JWTSecret() string
	JWTTTL()   time.Duration
}

type Auth interface {
    Register(ctx context.Context, login, email, password string) (int64, error)
    Login(ctx context.Context, login, password string) (int64, error)
}

type AuthService struct {
	repo UserRepo
	cfg  Config
}

func NewAuthService(r UserRepo, c Config) *AuthService { 
	return &AuthService{repo: r, cfg: c} 
}

func (s *AuthService) Register(ctx context.Context, login, email, password string) (int64, error) {
	// Проверка логина
	u, err := s.repo.ByLogin(ctx, login)
	if err != nil {
		return 0, err
	}
	if u != nil {
		return 0, ErrLoginTaken
	}

	// Проверка email
	u, err = s.repo.ByEmail(ctx, email)
	if err != nil {
		return 0, err
	}
	if u != nil {
		return 0, ErrEmailTaken
	}

	// Хеш пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// Создание пользователя
	return s.repo.Create(ctx, login, email, string(hash))
}
func (s *AuthService) Login(ctx context.Context, login, password string) (int64, error) {
	u, err := s.repo.ByLogin(ctx, login)
	if err != nil {
		return 0, err // ошибка БД / репозитория
	}
	if u == nil {
		return 0, errors.New("invalid credentials") // нет такого логина
	}

	// сверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return 0, errors.New("invalid credentials")
	}

	return u.ID, nil
	
}