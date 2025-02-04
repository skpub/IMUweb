package repository

import (
	"IMUbackend/db"
	entity "IMUbackend/internal/entity"
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

//
// Article = Markdown + Img
//

type IArticleRepository interface {
	Create(ctx context.Context, tx Tx, student string, imgs []*os.File, md entity.Markdown) (uuid.UUID, error)
}

type ArticleRepository struct {
	query *db.Queries
	minioClient *minio.Client
}

func NewArticleRepository(query *db.Queries, minioClient *minio.Client) IArticleRepository {
	return &ArticleRepository{query, minioClient}
}

func (a *ArticleRepository) Create(
	ctx context.Context,
	tx Tx,
	student string,
	imgs []*os.File,
	md entity.Markdown,
) (uuid.UUID, error) {
	id, err := tx.Queries().CreateMarkdown(ctx, db.CreateMarkdownParams{
		StudentID:   student,
		Title:       md.ArticleName,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	var imgIDs []uuid.UUID
	// 将来的にコイツもトランザクション的な何かで処理する
	for _, img := range imgs {
		path, err := tx.Queries().CreateImg(ctx, img.Name())
		if err != nil {
			return uuid.UUID{}, err
		}
		imgIDs = append(imgIDs, path)
	}
	for _, imgID := range imgIDs {
		err = tx.Queries().CreateMarkdownImgRel(ctx, db.CreateMarkdownImgRelParams{
			MarkdownID: id,
			ImgID:      imgID,
		})
		if err != nil {
			return uuid.UUID{}, err
		}
	}
	return id, nil
}
