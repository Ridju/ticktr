package user

import (
	"context"

	db "github.com/Ridju/ticktr/internal/db/sqlc"
)

type IUserRepository interface {
	CreateUser(CreateUserArgs, context.Context) (db.User, error)
	GetUserByMail(string, context.Context) (db.User, error)
	GetUserByID(int64, context.Context) (db.User, error)
}

type UserRepository struct {
	Store db.Store
}

func NewUserRepository(store db.Store) IUserRepository {
	return &UserRepository{
		Store: store,
	}
}

type CreateUserArgs struct {
	Username string
	Password string
	Email    string
}

func (r *UserRepository) CreateUser(args CreateUserArgs, ctx context.Context) (db.User, error) {
	arg := db.CreateUserParams{
		Username: args.Username,
		Password: args.Password,
		Email:    args.Email,
	}

	user, err := r.Store.CreateUser(ctx, arg)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByMail(email string, ctx context.Context) (db.User, error) {
	user, err := r.Store.GetUserByEmail(ctx, email)

	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(ID int64, ctx context.Context) (db.User, error) {
	user, err := r.Store.GetUserByID(ctx, ID)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
