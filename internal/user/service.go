package user

import (
	"context"
	"fmt"

	db "github.com/Ridju/ticktr/internal/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(username string, password string, email string, ctx context.Context) (db.User, error)
	LoginUser(email string, password string, ctx context.Context) (db.User, error)
	hashPassword(password string) (string, error)
	checkPassword(password string, hashedPassword string) error
}

type UserService struct {
	repo IUserRepository
}

func NewUserService(repo IUserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func (us *UserService) checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (us *UserService) CreateUser(username string, password string, email string, ctx context.Context) (db.User, error) {
	hashedPW, err := us.hashPassword(password)
	if err != nil {
		return db.User{}, err
	}

	args := CreateUserArgs{
		Username: username,
		Password: hashedPW,
		Email:    email,
	}

	user, err := us.repo.CreateUser(args, ctx)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (us *UserService) LoginUser(email string, password string, ctx context.Context) (db.User, error) {
	user, err := us.repo.GetUserByMail(email, ctx)
	if err != nil {
		return db.User{}, err
	}

	err = us.checkPassword(password, user.Password)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
