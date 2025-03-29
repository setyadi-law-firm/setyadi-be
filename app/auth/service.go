package auth

import (
	"errors"
)

type AuthService interface {
	Register(email, password, role string) (*User, error)
	Login(email, password string) (string, string, error)
}

type authService struct {
	repo AuthRepository
	util *Util
}

func NewAuthService(repo AuthRepository, util *Util) AuthService {
	return &authService{repo, util}
}

func (s *authService) Register(email, password, role string) (*User, error) {
	hashedPassword, err := s.util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}

	if err := s.repo.Save(user); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (s *authService) Login(email, password string) (string, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	if err := s.util.ComparePassword(user.Password, password); err != nil {
		return "", "", errors.New("invalid email or password")
	}

	accessToken, refreshToken, err := s.util.GenerateTokenPair(&User{Email: user.Email})
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
