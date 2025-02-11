package repository

import (
	"IMUbackend/db"
	"context"
	"database/sql"
	"fmt"
)

type IStudentRepository interface {
	Create(ctx context.Context, user db.Student) error
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
	return u.query.Login(ctx, db.LoginParams{
		ID:       id,
		Password: password,
	})
	// return tx.Queries().Login(ctx, db.LoginParams{
	// 	ID:       id,
	// 	Password: password,
	// })
}

func (u *StudentRepository) Create(ctx context.Context, user db.Student) error {
	_, err := u.query.FindStudentByID(ctx, user.ID)
	if err != nil {
		// err is not nil, but this means that the user does not exist
		// so we can create the user
		_, err := u.query.CreateStudent(ctx, db.CreateStudentParams{
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
