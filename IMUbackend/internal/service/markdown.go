package service

import (
	pb "IMUbackend/gen/imubackend"
	entity "IMUbackend/internal/domain"
	"time"
	"context"
)

func (s *IMUSrv) Create(ctx context.Context, p *pb.Markdown) error {
	md := entity.Markdown{
		ArticleName: p.ArticleName,
		Content:     p.Content,
		CreatedAt:   time.Now(),
	}

	err := s.md.Create(ctx, md)
	return err
}