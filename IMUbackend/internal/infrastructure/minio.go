package infrastructure

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient interface {
	PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error)
	GetObject(ctx context.Context, bucketName, objectName string, opts minio.GetObjectOptions) (*minio.Object, error)
}

func NewObjectStorageConnection(endpoint string, accessKeyID string, secret string) (*minio.Client, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secret, ""),
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

type minioMock struct {}

func NewMinioMock() MinioClient {
	return &minioMock{}
}

func (m *minioMock) PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error) {
	return minio.UploadInfo{}, nil
}

func (m *minioMock) GetObject(ctx context.Context, bucketName, objectName string, opts minio.GetObjectOptions) (*minio.Object, error) {
	return nil, nil
}
