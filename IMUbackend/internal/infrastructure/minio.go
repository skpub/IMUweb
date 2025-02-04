package infrastructure

import (
	"context"
	"io"
	entity "IMUbackend/internal/domain"
	repo "IMUbackend/internal/repository"

	"github.com/minio/minio-go/v7"
)

type IS3Client interface {
	Find(ctx context.Context, key string, bucket string) (io.Reader, error)
	Create(ctx context.Context, key string, payload io.Reader, bucket string) error
	Delete(ctx context.Context, key string, backet string) error
}

type MarkdownRepository struct {
	client *minio.Client
	bucket string
}

func NewMarkdownRepository(client *minio.Client, bucket string) repo.IMarkdownRepository {
	return &MarkdownRepository{client, bucket}
}

func (m *MarkdownRepository) Create(ctx context.Context, markdown entity.Markdown) error {
	// obj, _ := m.client.GetObject(ctx, m.bucket, markdown.ArticleName, minio.GetObjectOptions{})
	// if obj != nil {
		// return fmt.Errorf("already exists")
	// }
	// err := m.client.PutObject(ctx, m.bucket, markdown.ArticleName, markdown, m.bucket)
	// return err
	return nil
}

func (m *MarkdownRepository) Update(ctx context.Context, markdown entity.Markdown) error {
	return nil
}

func (m *MarkdownRepository) FindByID(ctx context.Context, articleName string) (entity.Markdown, error) {
	return entity.Markdown{}, nil
}

func (m *MarkdownRepository) Delete(ctx context.Context, articleName string) error {
	return nil
}