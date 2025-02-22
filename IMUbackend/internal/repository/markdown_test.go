package repository

import (
	"IMUbackend/db"
	entity "IMUbackend/internal/entity"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	mock "IMUbackend/internal/infrastructure/mock"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	mocker "github.com/stretchr/testify/mock"
)

func TestCreate_Success(t *testing.T) {
	qm := new(mock.Querier)
	minioClientMock := new(mock.MinioClient)

	articleRepo := NewArticleRepository(qm, minioClientMock, "")

	ctx := context.Background()

	
	t.Run("Create: Success (Create article with some imgs)", func(t *testing.T) {
		imgs := make([]*entity.NamedContent, 0)
		img := &entity.NamedContent{
			Name:    "test",
			Content: []byte("test"),
		}
		imgs = append(imgs, img)

		md := entity.Markdown{
			ArticleName: "test",
			Content:     "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		qm.On("CreateMarkdown", mocker.Anything, mocker.MatchedBy(func (p db.CreateMarkdownParams) bool {
			return p.Title == "test" && p.StudentID == "student" && strings.Contains(p.ContentPath, "student_test_")
		})).Return(uuid.New(), nil)
		qm.On("CreateImg", mocker.Anything, mocker.MatchedBy(func (imgName string) bool {
			return imgName == "test"
		})).Return(uuid.New(), nil)
		qm.On("CreateMarkdownImgRel", mocker.Anything, mocker.Anything).Return(nil)

		_, err := articleRepo.Create(ctx, "student", imgs, md)
		assert.NoError(t, err)
		qm.ExpectedCalls = nil
	})
	t.Run("Create: Success (Create article without imgs)", func(t *testing.T) {
		md := entity.Markdown{
			ArticleName: "test",
			Content:     "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		qm.On("CreateMarkdown", mocker.Anything, mocker.MatchedBy(func (p db.CreateMarkdownParams) bool {
			return p.Title == "test" && p.StudentID == "student" && strings.Contains(p.ContentPath, "student_test_")
		})).Return(uuid.New(), nil)

		_, err := articleRepo.Create(ctx, "student", nil, md)
		assert.NoError(t, err)
		qm.ExpectedCalls = nil
	})
	t.Run("Create: Fail (Can't create markdown)", func(t *testing.T) {
		md := entity.Markdown{
			ArticleName: "test",
			Content:     "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		qm.On("CreateMarkdown", mocker.Anything, mocker.Anything).Return(nil, fmt.Errorf("test"))
		//
		// No other calls expected.
		// 
		_, err := articleRepo.Create(ctx, "student", nil, md)
		assert.Error(t, err)
		qm.ExpectedCalls = nil
	})
}
