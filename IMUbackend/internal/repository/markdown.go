package repository

import (
	"IMUbackend/db"
	entity "IMUbackend/internal/entity"
	"IMUbackend/internal/infrastructure"
	"context"
	"io"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

//
// Article = Markdown + Img
//

type IArticleRepository interface {
	Create(ctx context.Context, student string, imgs []*entity.NamedContent, md entity.Markdown) (uuid.UUID, error)
	ListAll(ctx context.Context) ([]db.ListMarkdownRow, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Article, error)
}

type ArticleRepository struct {
	query       db.Querier
	minioClient infrastructure.MinioClient
	bucket      string
}

func NewArticleRepository(query db.Querier, minioClient infrastructure.MinioClient, bucket string) IArticleRepository {
	return &ArticleRepository{query, minioClient, bucket}
}

func (a *ArticleRepository) Create(
	ctx context.Context,
	student string,
	imgs []*entity.NamedContent,
	md entity.Markdown,
) (uuid.UUID, error) {
	// db

	contentPath := student + "_" + md.ArticleName + "_" + time.Now().String() + ".md"
	markdownID, err := a.query.CreateMarkdown(ctx, db.CreateMarkdownParams{
		StudentID:   student,
		Title:       md.ArticleName,
		ContentPath: contentPath,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	// minio
	rContent := strings.NewReader(md.Content)
	_, err = a.minioClient.PutObject(ctx, a.bucket, contentPath, rContent, -1, minio.PutObjectOptions{})
	if err != nil {
		return uuid.UUID{}, err
	}

	var imgIDs []uuid.UUID
	for _, img := range imgs {
		// db
		imgPath, err := a.query.CreateImg(ctx, img.Name)
		if err != nil {
			return uuid.UUID{}, err
		}
		imgIDs = append(imgIDs, imgPath)

		err = a.query.CreateMarkdownImgRel(ctx, db.CreateMarkdownImgRelParams{
			MarkdownID: markdownID,
			ImgID:      imgPath,
		})
		if err != nil {
			return uuid.UUID{}, err
		}

		// minio
		// 将来的にコイツもトランザクション的な何かで処理する
		imgRdr := strings.NewReader(md.Content)
		_, err = a.minioClient.PutObject(ctx, a.bucket, imgPath.String(), imgRdr, -1, minio.PutObjectOptions{})
		if err != nil {
			return uuid.UUID{}, err
		}
	}
	return markdownID, nil
}

func (a *ArticleRepository) ListAll(ctx context.Context) ([]db.ListMarkdownRow, error) {
	id, err := a.query.ListMarkdown(ctx)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (a *ArticleRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Article, error) {
	article := &entity.Article{}
	r, err := a.query.GetArticle(ctx, id)
	if err != nil {
		return &entity.Article{}, err
	}
	if len(r) == 0 {
		return &entity.Article{}, nil
	}
	article.ID = r[0].ID.String()
	article.StudentID = r[0].StudentID
	article.Title = r[0].Title
	mdFile, err := a.minioClient.GetObject(ctx, a.bucket, r[0].ContentPath, minio.GetObjectOptions{})
	if err != nil {
		return &entity.Article{}, err
	}
	md, err := io.ReadAll(mdFile)
	if err != nil {
		return &entity.Article{}, err
	}
	article.Content = string(md)
	article.CreatedAt = r[0].Since
	article.UpdatedAt = r[0].Updated
	for _, record := range r {
		imgPath := record.ImgID
		img, err := a.minioClient.GetObject(ctx, a.bucket, imgPath.String(), minio.GetObjectOptions{})
		if err != nil {
			return &entity.Article{}, err
		}
		imgBytes, err := io.ReadAll(img)
		if err != nil {
			return &entity.Article{}, err
		}
		article.Imgs = append(article.Imgs, &entity.NamedContent{
			Name:    imgPath.String(),
			Content: imgBytes,
		})
	}
	return article, nil
}
