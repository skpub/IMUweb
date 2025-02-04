package repository

import (
	"IMUbackend/db"
	"context"
	"database/sql"
	"fmt"
)

type IUserRepository interface {
	Create(ctx context.Context, tx Tx, user db.User) error
	UpdateBio(ctx context.Context, tx Tx, id string, bio string) error
	FindByID(ctx context.Context, tx Tx, id string) (db.User, error)
	Login(ctx context.Context, tx Tx, id string, password string) error
	Delete(ctx context.Context, tx Tx, id string) error
}

type UserRepository struct {
	query *db.Queries
}

func NewUserRepository(query *db.Queries) IUserRepository {
	return &UserRepository{query}
}

func (u *UserRepository) Login(ctx context.Context,tx Tx, id string, password string) error {
	return tx.Queries().Login(ctx, db.LoginParams{
		ID: id,
		Password: password,
	})
}

func (u *UserRepository) Create(ctx context.Context, tx Tx ,user db.User) error {
	_, err := tx.Queries().FindUserByID(ctx, user.ID)
	if err != nil {
		// err is not nil, but this means that the user does not exist
		// so we can create the user
		tx.Queries().CreateUser(ctx, db.CreateUserParams{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			Password: user.Password,
		})
		return nil
	}
	// Already exists, so return error
	return fmt.Errorf("already exists")
}

func (u *UserRepository) UpdateBio(ctx context.Context, tx Tx, id string, bio string) error {
	_, err := tx.Queries().FindUserByID(ctx, id)
	if err != nil {
		// Not found
		return err
	}
	return tx.Queries().UpdateUserBio(ctx, db.UpdateUserBioParams{
		ID: id,
		Bio: sql.NullString{ String: bio, Valid: true },
	})
}

func (u *UserRepository) FindByID(ctx context.Context, tx Tx, articleName string) (db.User, error) {
	return db.User{}, nil
}

func (u *UserRepository) Delete(ctx context.Context, tx Tx, articleName string) error {
	return nil
}
