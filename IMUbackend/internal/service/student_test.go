package service

import (
	dbb "IMUbackend/db"
	pb "IMUbackend/gen/imubackend"
	dbMock "IMUbackend/internal/infrastructure/mock"
	repoMock "IMUbackend/internal/repository/mock"
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStudent(t *testing.T) {
	userId := "testuser"
	userName := "testusername"
	password := "testpwd"
	email := "test@testmail.com"
	salt := "testsalt"
	hash, err := genhash(password, salt)

	student := new(repoMock.IStudentRepository)
	article := new(repoMock.IArticleRepository)
	db := new(dbMock.IDBTX)
	srv := NewIMUSrv(article, student, "testsec", salt, db)
	ctx := context.Background()
	t.Run("Login: Success", func(t *testing.T) {
		assert.NoError(t, err)
		student.On("Login", ctx, userId, hash).Return(nil)
		arg := &pb.Login2{
			StudentID: &userId,
			Password:  &password,
		}
		_, err = srv.Login(ctx, arg)
		assert.NoError(t, err)
		student.ExpectedCalls = nil
	})
	t.Run("CreateStudent: Success", func(t *testing.T) {
		student.On("Create", ctx, dbb.Student{
			ID:       userId,
			Name:     userName,
			Email:    email,
			Password: hash,
		}).Return(nil)
		db.On("BeginTx", ctx, (*sql.TxOptions)(nil)).Return(&sql.Tx{}, nil)
		db.On("Commit", mock.Anything).Return(nil)
		_, err := srv.Signup(ctx, &pb.SignupPayload{
			Name: &userName,
			StudentID:   &userId,
			Email:       &email,
			Password:    &password,
		})
		assert.NoError(t, err)
		student.ExpectedCalls = nil
		db.ExpectedCalls = nil
	})
	t.Run("CreateStudent: Fail (Rollback)", func(t *testing.T) {
		student.On("Create", ctx, dbb.Student{
			ID:       userId,
			Name:     userName,
			Email:    email,
			Password: hash,
		}).Return(fmt.Errorf("クエリのどっかでエラー"))
		db.On("BeginTx", ctx, (*sql.TxOptions)(nil)).Return(&sql.Tx{}, nil)
		db.On("Rollback", mock.Anything).Return(nil)
		_, err := srv.Signup(ctx, &pb.SignupPayload {
			Name: &userName,
			StudentID:   &userId,
			Email:       &email,
			Password:    &password,
		})
		assert.Error(t, err)
		student.ExpectedCalls = nil
		db.ExpectedCalls = nil
	})
}
