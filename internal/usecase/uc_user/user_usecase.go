package uc_user

import (
	"errors"

	user "github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/domain/d_user"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type PasswordHash interface {
	Hash(password string) (string, error)
	Compare(hashed, plain string) bool
}

type JwtGenerator interface {
	GenerateToken(userID int64, email string) (string, error)
}

type UserUseCase struct {
	repo   user.Repository
	jwt    JwtGenerator
	hasher PasswordHash
}

func NewUserUseCase(repo user.Repository, jwt JwtGenerator, hasher PasswordHash) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		jwt:    jwt,
		hasher: hasher,
	}
}

func (u *UserUseCase) Register(name, email, password string) (*user.UserE, error) {
	existing, _ := u.repo.GetByEmail(email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := u.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	user := &user.UserE{
		Name:     name,
		Email:    email,
		Password: hashed,
	}

	if err := u.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) Login(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil || user == nil {
		return "", ErrInvalidCredentials
	}

	if !u.hasher.Compare(user.Password, password) {
		return "", ErrInvalidCredentials
	}

	return u.jwt.GenerateToken(user.ID, user.Email)
}

func (u *UserUseCase) GetProfile(userID int64) (*user.UserE, error) {
	return u.repo.GetByID(userID)
}
