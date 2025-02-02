package infrastructure

import (
	"context"
	"io"
	"os"
    "strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type IS3Client interface {
	Find(context.Context, string) (io.Reader, error)
	Create(context.Context, string, io.Reader) (error)
	Delete(context.Context, string) (error)
}

type S3Client struct {
	client *minio.Client
	bucket string
}

func NewS3Client() (IS3Client, error) {
	endpoint := os.Getenv("MINIO_SERVER_URL")
	accessKeyID := os.Getenv("MINIO_ROOT_USER")
	secret := os.Getenv("MINIO_ROOT_PASSWORD")
	bucket := os.Getenv("MDBUCKET")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secret, ""),
	})
	client_wrapper := S3Client{client: client, bucket: bucket}
	return client_wrapper, err
}

func (client S3Client) Find(ctx context.Context, key string) (io.Reader, error) {
	content, err := client.client.GetObject(ctx, client.bucket, key, minio.GetObjectOptions{})
	return content, err
}

func (client S3Client) Create(ctx context.Context, key string, content io.Reader) error {
	var r io.Reader = strings.NewReader("test")
	_, err := client.client.PutObject(ctx, client.bucket, key, r, -1, minio.PutObjectOptions{})
	return err
}

func (client S3Client) Delete(ctx context.Context, key string) error {
	err := client.client.RemoveObject(ctx, client.bucket, key, minio.RemoveObjectOptions{})
	return err
}
