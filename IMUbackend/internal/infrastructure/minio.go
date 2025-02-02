package infrastructure

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

type IS3Client interface {
	Find(ctx context.Context, key string, bucket string) (io.Reader, error)
	Create(ctx context.Context, key string, payload io.Reader, bucket string) error
	Delete(ctx context.Context, key string, backet string) error
}

type S3Client struct {
	client *minio.Client
}

func NewS3Client(client *minio.Client) IS3Client {
	return S3Client{client: client}
}

func (client S3Client) Find(ctx context.Context, key string, bucket string) (io.Reader, error) {
	content, err := client.client.GetObject(ctx, bucket, key, minio.GetObjectOptions{})
	return content, err
}

func (client S3Client) Create(ctx context.Context, key string, content io.Reader, bucket string) error {
	_, err := client.client.PutObject(ctx, bucket, key, content, -1, minio.PutObjectOptions{})
	return err
}

func (client S3Client) Delete(ctx context.Context, key string, bucket string) error {
	err := client.client.RemoveObject(ctx, bucket, key, minio.RemoveObjectOptions{})
	return err
}
