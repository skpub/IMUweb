package repository

import (
	"IMUbackend/db"
	"context"
	"database/sql"
	"fmt"
)

type IStudentRepository interface {
	Create(ctx context.Context, tx Tx, user db.Student) error
	UpdateBio(ctx context.Context, tx Tx, id string, bio string) error
	FindByID(ctx context.Context, tx Tx, id string) (db.Student, error)
	Login(ctx context.Context, tx Tx, id string, password string) error
	Delete(ctx context.Context, tx Tx, id string) error
}

type StudentRepository struct {
	query *db.Queries
}

func NewStudentRepository(query *db.Queries) IStudentRepository {
	return &StudentRepository{query}
}

func (u *StudentRepository) Login(ctx context.Context, tx Tx, id string, password string) error {
	return tx.Queries().Login(ctx, db.LoginParams{
		ID:       id,
		Password: password,
	})
}

func (u *StudentRepository) Create(ctx context.Context, tx Tx, user db.Student) error {
	_, err := tx.Queries().FindStudentByID(ctx, user.ID)
	if err != nil {
		// err is not nil, but this means that the user does not exist
		// so we can create the user
		_, err := tx.Queries().CreateStudent(ctx, db.CreateStudentParams{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
		return err
	}
	// Already exists, so return error
	return fmt.Errorf("already exists")
}

func (u *StudentRepository) UpdateBio(ctx context.Context, tx Tx, id string, bio string) error {
	_, err := tx.Queries().FindStudentByID(ctx, id)
	if err != nil {
		// Not found
		return err
	}
	return tx.Queries().UpdateStudentBio(ctx, db.UpdateStudentBioParams{
		ID:  id,
		Bio: sql.NullString{String: bio, Valid: true},
	})
}

func (u *StudentRepository) FindByID(ctx context.Context, tx Tx, articleName string) (db.Student, error) {
	return db.Student{}, nil
}

func (u *StudentRepository) Delete(ctx context.Context, tx Tx, articleName string) error {
	return nil
}
