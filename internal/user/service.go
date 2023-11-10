package user

import (
	"fmt"

	"github.com/Ridju/ticktr/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(string, string, string) (db.User, error)
	LoginUser(string, string) (db.User, error)
	hashPassword(string) (string, error)
	checkPassword(string, string) error
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

func (us *UserService) CreateUser(username string, password string, email string) (db.User, error) {
	hashedPW, err := us.hashPassword(password)
	if err != nil {
		return db.User{}, err
	}

	user, err := us.repo.CreateUser(username, hashedPW, email)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (us *UserService) LoginUser(email string, password string) (db.User, error) {
	user, err := us.repo.GetUserByMail(email)
	if err != nil {
		return db.User{}, err
	}

	err = us.checkPassword(password, user.Password)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
