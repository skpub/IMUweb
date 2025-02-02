package service

import (
	pb "IMUbackend/gen/imubackend"
	entity "IMUbackend/internal/domain"
	repo "IMUbackend/internal/repository"
	"context"
	"time"
)

type MarkdownSrv struct {
	repo repo.IMarkdownRepository
}

func NewMarkdownService(repo repo.IMarkdownRepository) pb.Service {
	return &MarkdownSrv{repo}
}

func (s *MarkdownSrv) Create(ctx context.Context, p *pb.Markdown) error {
	md := entity.Markdown{
		ArticleName: p.ArticleName,
		Content:     p.Content,
		CreatedAt:   time.Now(),
	}

	err := s.repo.Create(ctx, md)
	return err
}
