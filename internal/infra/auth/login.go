package auth

import (
	"errors"

	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type Login struct {
	userQueryRepository entities.UserQueryRepository
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewLogin(userQryRepo entities.UserQueryRepository) *Login {
	return &Login{
		userQueryRepository: userQryRepo,
	}
}

func (l *Login) Authenticate(userCredentials UserCredentials) (*TokenResponse, error) {
	user, err := l.userQueryRepository.FindByEmail(userCredentials.Email)
	if err != nil || user.Id == 0 {
		return nil, errors.New("invalid credentials")
	}

	if err := user.CheckPassword(userCredentials.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := GenerateToken(user.Email, user.FirstName, user.LastName, user.Id)
	var tokenResponse TokenResponse
	tokenResponse.Token = token
	return &tokenResponse, err
}
