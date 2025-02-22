package repository

import (
	"IMUbackend/db"
	mock "IMUbackend/internal/infrastructure/mock"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	mocker "github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	qm := new(mock.Querier)
	mm := new(mock.MinioClient)
	studentRepo := NewStudentRepository(qm, mm, "bucket")
	ctx := context.Background()

	t.Run("CreateStudent: Success", func(t *testing.T) {
		id := "testuser"
		student := db.CreateStudentParams {
			ID: id,
			Name: "test",
			Email: "test@testmail.com",
			Password: "test",
		}
		qm.On("FindStudentByID", ctx, id).Return(db.Student{}, fmt.Errorf("not found"))
		qm.On("CreateStudent", ctx, mocker.Anything).Return("user_id", nil)
		_, err := studentRepo.Create(ctx, student)
		assert.NoError(t, err)
		qm.ExpectedCalls = nil
	})
	t.Run("CreateStudent: Fail (Already exists)", func(t *testing.T) {
		id := "testuser"
		student := db.CreateStudentParams {
			ID: id,
			Name: "test",
			Email: "test@testmail.com",
			Password: "test",
		}
		qm.On("FindStudentByID", ctx, id).Return(db.Student{}, nil)
		// CreateStudentは呼ばれないはず。
		_, err := studentRepo.Create(ctx, student)
		assert.Error(t, err)
		qm.ExpectedCalls = nil
	})
}
