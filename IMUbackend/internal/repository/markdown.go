package repository

import (
	entity "IMUbackend/internal/domain"
	"context"
)

type IMarkdownRepository interface {
	Create(ctx context.Context, markdown entity.Markdown) error
	Update(ctx context.Context, markdown entity.Markdown) error
	FindByID(ctx context.Context, articleName string) (entity.Markdown, error)
	Delete(ctx context.Context, articleName string) error
}
