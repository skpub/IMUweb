package repository

import (
	"IMUbackend/db"
	"IMUbackend/internal/infrastructure"
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/minio/minio-go/v7"
)

type IStudentRepository interface {
	Create(ctx context.Context, user db.CreateStudentParams) (string, error)
	FindByID(ctx context.Context, id string) (db.Student, error)
	GetProfile(ctx context.Context) (StudentProfile, error)
	GetProfiles(ctx context.Context) ([]*StudentProfile, error)
	UpdateBio(ctx context.Context, id string, bio string) error
	UpdateImg(ctx context.Context, id string, img []byte) error
	Login(ctx context.Context, id string, password string) error
	Delete(ctx context.Context, id string) error
}

type StudentRepository struct {
	query       db.Querier
	minioClient infrastructure.MinioClient
	bucket      string
}

func NewStudentRepository(query db.Querier, minioClient infrastructure.MinioClient, bucket string) IStudentRepository {
	return &StudentRepository{query, minioClient, bucket}
}

func (u *StudentRepository) Login(ctx context.Context, id string, password string) error {
	num, err := u.query.Login(ctx, db.LoginParams{
		ID:       id,
		Password: password,
	})
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("auth error")
	}
	return nil
	// return tx.Queries().Login(ctx, db.LoginParams{
	// 	ID:       id,
	// 	Password: password,
	// })
}

func (u *StudentRepository) Create(ctx context.Context, user db.CreateStudentParams) (string, error) {
	_, err := u.query.FindStudentByID(ctx, user.ID)
	if err != nil {
		// err is not nil, but this means that the user does not exist
		// so we can create the user
		id, err := u.query.CreateStudent(ctx, user)
		return id, err
	}
	// Already exists, so return error
	return "", fmt.Errorf("already exists")
}

func (u *StudentRepository) UpdateBio(ctx context.Context, id string, bio string) error {
	_, err := u.query.FindStudentByID(ctx, id)
	if err != nil {
		// Not found
		return err
	}
	return u.query.UpdateStudentBio(ctx, db.UpdateStudentBioParams{
		ID:  id,
		Bio: sql.NullString{String: bio, Valid: true},
	})
}

func (u *StudentRepository) UpdateImg(ctx context.Context, id string, img []byte) error {
	_, err := u.query.FindStudentByID(ctx, id)
	if err != nil {
		// Not found
		return err
	}
	imgId, err := u.query.UpdateStudentImg(ctx, id)
	if err != nil {
		return err
	}
	imgIdString := imgId.UUID.String()
	imgRdr := bytes.NewReader(img)
	_, err = u.minioClient.PutObject(ctx, u.bucket, imgIdString, imgRdr, -1, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (u *StudentRepository) FindByID(ctx context.Context, studentId string) (db.Student, error) {
	return u.query.FindStudentByID(ctx, studentId)
}

type StudentProfile struct {
	StudentID *string
	Name      *string
	Bio       *string
	Img       *[]byte
}

func (u *StudentRepository) GetProfile(ctx context.Context) (StudentProfile, error) {
	student, err := u.query.FindStudentByID(ctx, ctx.Value("studentId").(string))
	if err != nil {
		return StudentProfile{}, err
	}

	studentProfile := StudentProfile {
		StudentID: &student.ID,
		Name: &student.Name,
		Bio: &student.Bio.String,
	}

	if student.ImgPath.Valid {
		img, err := u.minioClient.GetObject(ctx, u.bucket, student.ImgPath.UUID.String(), minio.GetObjectOptions{})
		if err != nil {
			return StudentProfile{}, err
		}
		imgBytes := make([]byte, 0)
		_, err = img.Read(imgBytes)
		if err != nil {
			return StudentProfile{}, err
		}

		studentProfile.Img = &imgBytes
	}
	return studentProfile, nil
}

func (u *StudentRepository) GetProfiles(ctx context.Context) ([]*StudentProfile, error) {
	students, err := u.query.FindStudents(ctx)
	if err != nil {
		return nil, err
	}

	studentProfiles := make([]*StudentProfile, 0)

	for _, student := range students {
		studentProfile := StudentProfile {
			StudentID: &student.ID,
			Name: &student.Name,
			Bio: &student.Bio.String,
		}
		if student.ImgPath.Valid {
			img, err := u.minioClient.GetObject(ctx, u.bucket, student.ImgPath.UUID.String(), minio.GetObjectOptions{})
			if err != nil {
				return nil, err
			}
			imgBytes := make([]byte, 0)
			_, err = img.Read(imgBytes)
			if err != nil {
				return nil, err
			}
			studentProfile.Img = &imgBytes
		}
		studentProfiles = append(studentProfiles, &studentProfile)
	}
	return studentProfiles, nil
}

func (u *StudentRepository) Delete(ctx context.Context, studentId string) error {
	student, err := u.query.FindStudentByID(ctx, studentId)
	if err != nil {
		return err
	}
	if student.ImgPath.Valid {
		err := u.minioClient.RemoveObject(ctx, u.bucket, student.ImgPath.UUID.String(), minio.RemoveObjectOptions{})
		if err != nil {
			return err
		}
	}
	return u.query.DeleteStudent(ctx, studentId)
}
