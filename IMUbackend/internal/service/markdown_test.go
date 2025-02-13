package service

import (
	entity "IMUbackend/internal/entity"
	dbMock "IMUbackend/internal/infrastructure/mock"
	repoMock "IMUbackend/internal/repository/mock"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	pb "IMUbackend/gen/imubackend"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListArticle(t *testing.T) {
	student := new(repoMock.IStudentRepository)
	article := new(repoMock.IArticleRepository)
	db := new(dbMock.IDBTX)
	srv := NewIMUSrv(article, student, "testsec", "testsalt", db)
	t.Run("ListArticle: Success", func(t *testing.T) {
		ctx := context.Background()
		retList := []uuid.UUID{uuid.New(), uuid.New()}
		article.On("ListAll", ctx).Return(retList, nil)
		result, err := srv.ListArticle(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		article.ExpectedCalls = nil
	})
	t.Run("ListArticle: Fail", func(t *testing.T) {
		ctx := context.Background()
		article.On("ListAll", ctx).Return(nil, fmt.Errorf("test"))
		result, err := srv.ListArticle(ctx)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestArticle(t *testing.T) {
	student := new(repoMock.IStudentRepository)
	article := new(repoMock.IArticleRepository)
	db := new(dbMock.IDBTX)
	srv := NewIMUSrv(article, student, "testsec", "testsalt", db)

	t.Run("GetArticle: Success", func(t *testing.T) {
		ctx := context.Background()
		argUUID := uuid.New()
		argUUIDStr := argUUID.String()
		article.On("FindByID", ctx, argUUID).Return(&entity.Article{}, nil)
		result, err := srv.GetArticle(ctx, argUUIDStr)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		article.ExpectedCalls = nil
	})
	t.Run("GetArticle: Fail", func(t *testing.T) {
		ctx := context.Background()
		argUUID := uuid.New()
		argUUIDStr := argUUID.String()
		article.On("FindByID", ctx, argUUID).Return(nil, fmt.Errorf("test"))
		result, err := srv.GetArticle(ctx, argUUIDStr)
		assert.Error(t, err)
		assert.Nil(t, result)
	}) 
	t.Run("CreateArticle: Success (with some imgs)", func(t *testing.T) {
		ctx := context.Background()
		md := entity.Markdown {
			ArticleName: "testArticle",
			Content: "testArticleContent",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// start: mock setup
		argImgs := make([]*entity.NamedContent, 0)
		argImgs = append(argImgs, &entity.NamedContent{
			Name: "testImg",
			Content: []byte("testImgContent"),
		})
		ctx = context.WithValue(ctx, "studentId", "testuser")
		article.On("Create", ctx, "testuser", argImgs, mock.Anything).Return(uuid.New(), nil)
		db.On("BeginTx", ctx, (*sql.TxOptions)(nil)).Return(&sql.Tx{}, nil)
		db.On("Commit", mock.Anything).Return(nil)
		// コミットが起こるのであって、ロールバックは起こらない。
		// end: mock setup
		imgs := make([]*pb.File, 0)
		imgName := "testImg"
		imgs = append(imgs, &pb.File{
			Name: &imgName,
			Content: []byte("testImgContent"),
		})
		arg := pb.CreateArticlePayload{
			ArticleName: md.ArticleName,
			Content: md.Content,
			Image: imgs,
			Token: "testToken",
		}
		// call
		err := srv.CreateArticle(ctx, &arg)

		assert.NoError(t, err)
		db.ExpectedCalls = nil
		article.ExpectedCalls = nil
	})
	t.Run("CreateArticle: Success (without imgs)", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "studentId", "testuser")
		emptylist := make([]*entity.NamedContent, 0)
		article.On("Create", ctx, "testuser", emptylist, mock.Anything).Return(uuid.New(), nil)
		db.On("BeginTx", ctx, (*sql.TxOptions)(nil)).Return(&sql.Tx{}, nil)
		db.On("Commit", mock.Anything).Return(nil)

		// call
		err := srv.CreateArticle(ctx, &pb.CreateArticlePayload{
			ArticleName: "testArticle",
			Content: "testArticleContent",
			Image: nil,
			Token: "testToken",
		})
		assert.NoError(t, err)
		article.ExpectedCalls = nil
		db.ExpectedCalls = nil
	})
	t.Run("CreateArticle: Fail (Can't create article)", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "studentId", "testuser")
		emptylist := make([]*entity.NamedContent, 0)
		article.On("Create", ctx, "testuser", emptylist, mock.Anything).Return(nil, fmt.Errorf("test"))
		db.On("BeginTx", ctx, (*sql.TxOptions)(nil)).Return(&sql.Tx{}, nil)
		db.On("Rollback", mock.Anything).Return(nil)
		// ロールバックが起こるのであって、コミットは起こらない。

		// call
		err := srv.CreateArticle(ctx, &pb.CreateArticlePayload{
			ArticleName: "testArticle",
			Content: "testArticleContent",
			Image: nil,
			Token: "testToken",
		})
		assert.Error(t, err)
		article.ExpectedCalls = nil
		db.ExpectedCalls = nil
	})
}
