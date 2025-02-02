package repository

import (
	entity "IMUbackend/internal/domain"
	infra "IMUbackend/internal/infrastructure"
	"context"
	"fmt"
)

type IMarkdownRepository interface {
	Create(ctx context.Context, markdown entity.Markdown) error
	Update(ctx context.Context, markdown entity.Markdown) error
	FindByID(ctx context.Context, articleName string) (entity.Markdown, error)
	Delete(ctx context.Context, articleName string) error
}

type MarkdownRepository struct {
	client infra.IS3Client
	bucket string
}

func NewMarkdownRepository(client infra.IS3Client, bucket string) IMarkdownRepository {
	return &MarkdownRepository{client, bucket}
}

func (m *MarkdownRepository) Create(ctx context.Context, markdown entity.Markdown) error {
	obj, _ := m.client.Find(ctx, markdown.ArticleName, m.bucket)
	if obj != nil {
		return fmt.Errorf("already exists")
	}
	err := m.client.Create(ctx, markdown.ArticleName, markdown, m.bucket)
	return err
}

func (m *MarkdownRepository) Update(ctx context.Context, markdown entity.Markdown) error {
	content, err := m.client.Find(ctx, markdown.ArticleName, m.bucket)
	if err == nil {
		return fmt.Errorf("not found")
	}
	contentmd := content.(entity.Markdown)
	if contentmd == markdown {
		return fmt.Errorf("no changes")
	}

	err = m.client.Create(ctx, markdown.ArticleName, markdown, m.bucket)
	return err
}

func (m *MarkdownRepository) FindByID(ctx context.Context, articleName string) (entity.Markdown, error) {
	content, err := m.client.Find(ctx, articleName, m.bucket)
	if err != nil {
		return entity.Markdown{}, err
	}

	contentmd, ok := content.(entity.Markdown)
	if !ok {
		return entity.Markdown{}, fmt.Errorf("data corrupted")
	}

	return contentmd, nil
}

func (m *MarkdownRepository) Delete(ctx context.Context, articleName string) error {
	return m.client.Delete(ctx, articleName, m.bucket)
}
