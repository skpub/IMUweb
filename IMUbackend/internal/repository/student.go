package repository

import (
	"IMUbackend/db"
	"context"
	"database/sql"
	"fmt"
)

type IStudentRepository interface {
	Create(ctx context.Context, user db.CreateStudentParams) (string, error)
	UpdateBio(ctx context.Context, id string, bio string) error
	FindByID(ctx context.Context, id string) (db.Student, error)
	Login(ctx context.Context, id string, password string) error
	Delete(ctx context.Context, id string) error
}

type StudentRepository struct {
	query db.Querier
}

func NewStudentRepository(query db.Querier) IStudentRepository {
	return &StudentRepository{query}
}

func (u *StudentRepository) Login(ctx context.Context, id string, password string) error {
	num, err := u.query.Login(ctx, db.LoginParams{
		ID:       id,
		Password: password,
	})
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("auth error")
	}
	return nil
	// return tx.Queries().Login(ctx, db.LoginParams{
	// 	ID:       id,
	// 	Password: password,
	// })
}

func (u *StudentRepository) Create(ctx context.Context, user db.CreateStudentParams) (string, error) {
	_, err := u.query.FindStudentByID(ctx, user.ID)
	if err != nil {
		// err is not nil, but this means that the user does not exist
		// so we can create the user
		id, err := u.query.CreateStudent(ctx, user)
		return id, err
	}
	// Already exists, so return error
	return "", fmt.Errorf("already exists")
}

func (u *StudentRepository) UpdateBio(ctx context.Context, id string, bio string) error {
	_, err := u.query.FindStudentByID(ctx, id)
	if err != nil {
		// Not found
		return err
	}
	return u.query.UpdateStudentBio(ctx, db.UpdateStudentBioParams{
		ID:  id,
		Bio: sql.NullString{String: bio, Valid: true},
	})
}

func (u *StudentRepository) FindByID(ctx context.Context, studentId string) (db.Student, error) {
	return u.query.FindStudentByID(ctx, studentId)
}

func (u *StudentRepository) Delete(ctx context.Context, studentId string) error {
	return u.query.DeleteStudent(ctx, studentId)
}
