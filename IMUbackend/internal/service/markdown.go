package service

import (
	pb "IMUbackend/gen/imubackend"
	entity "IMUbackend/internal/entity"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (s *IMUSrv) ListArticle(ctx context.Context) (*pb.ListArticleResult, error) {
	idsUUID, err := s.article.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	ids := pb.ListArticleResult{}
	for _, id := range idsUUID {
		idStr := id.ID.String()
		updated := id.Updated.String()
		ids.List = append(ids.List, &pb.ArticleIDName{
			ID:      &idStr,
			Name:    &id.Title,
			Updated: &updated,
		})
	}
	return &ids, nil
}

func (s *IMUSrv) GetArticle(ctx context.Context, p *pb.GetArticlePayload) (*pb.GetArticleResult, error) {
	id, err := uuid.Parse(*p.ID)
	if err != nil {
		return nil, err
	}
	article, err := s.article.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	imgFiles := make([]*pb.File, 0)
	for _, img := range article.Imgs {
		imgFileName := new(string)
		*imgFileName = "test"
		imgFile := &pb.File{
			Name:    imgFileName,
			Content: img.Content,
		}
		imgFiles = append(imgFiles, imgFile)
	}
	created := article.CreatedAt.String()
	updated := article.UpdatedAt.String()
	articleRet := &pb.GetArticleResult{
		ID:          &article.ID,
		StudentID:   &article.StudentID,
		ArticleName: &article.Title,
		Content:     &article.Content,
		Image:       imgFiles,
		CreatedAt:   &created,
		UpdatedAt:   &updated,
	}
	return articleRet, nil
}

func (s *IMUSrv) CreateArticle(ctx context.Context, p *pb.CreateArticlePayload) error {
	studentId, ok := ctx.Value("studentId").(string)
	if !ok {
		return fmt.Errorf("failed to get studentId. set studentId in context")
	}
	now := time.Now()
	md := entity.Markdown{
		ArticleName: p.ArticleName,
		Content:     p.Content,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tx, err := s.db.BeginTx(ctx, nil)

	defer func() {
		if err != nil {
			s.db.Rollback(tx)
		} else {
			s.db.Commit(tx)
		}
	}()

	imgs := make([]*entity.NamedContent, 0)
	for _, pImg := range p.Image {
		img := &entity.NamedContent{
			Name:    *pImg.Name,
			Content: pImg.Content,
		}
		imgs = append(imgs, img)
	}
	_, err = s.article.Create(ctx, studentId, imgs, md)
	return err
}
