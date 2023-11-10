package user

import (
	"github.com/Ridju/ticktr/internal/db"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(string, string, string) (db.User, error)
	GetUserByMail(string) (db.User, error)
	GetUserByID(int) (db.User, error)
}

type GORMRepository struct {
	db *gorm.DB
}

func NewGORMRepository(db *gorm.DB) IUserRepository {
	return &GORMRepository{
		db: db,
	}
}

func (r *GORMRepository) CreateUser(username string, password string, email string) (db.User, error) {

	user := db.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	result := r.db.Create(&user)
	if result.Error != nil {
		return db.User{}, result.Error
	}

	return user, nil
}

func (r *GORMRepository) GetUserByMail(email string) (db.User, error) {
	var user db.User
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return db.User{}, result.Error
	}

	return user, nil
}

func (r *GORMRepository) GetUserByID(ID int) (db.User, error) {
	var user db.User
	result := r.db.First(&user, "id = ?", ID)
	if result.Error != nil {
		return db.User{}, result.Error
	}

	return user, nil
}
