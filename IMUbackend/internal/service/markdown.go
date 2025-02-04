package service

import (
	pb "IMUbackend/gen/imubackend"
	entity "IMUbackend/internal/entity"
	"context"
	"os"
	"time"
)

func (s *IMUSrv) CreateMarkdown(ctx context.Context, p *pb.CreateMarkdownAttr) error {
	md := entity.Markdown{
		ArticleName: p.ArticleName,
		Content:     p.Content,
		CreatedAt:   time.Now(),
	}

	tx, err := s.tx.BeginTx(ctx)

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var files []*os.File

	for _, content := range p.Image {
		tmpFile, err := os.CreateTemp("", "tmp")
		if err != nil {
			return err
		}
		if _, err := tmpFile.Write(content); err != nil {
			return err
		}
		tmpFile.Close()
		file, err := os.Open(tmpFile.Name())
		if err != nil {
			return err
		}
		files = append(files, file)
	}

	_, err = s.article.Create(ctx, tx, *p.StudentID, files, md)
	return err
}
