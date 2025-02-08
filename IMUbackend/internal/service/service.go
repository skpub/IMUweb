package service

import (
	pb "IMUbackend/gen/imubackend"
	repo "IMUbackend/internal/repository"
	"database/sql"
)

type IMUSrv struct {
	article   repo.IArticleRepository
	user      repo.IStudentRepository
	jwtsecret string
	salt      string
	db        *sql.DB
}


func NewIMUSrv(article repo.IArticleRepository, user repo.IStudentRepository, jwtsecret string, salt string, db *sql.DB) pb.Service {
	return &IMUSrv{
		article,
		user,
		jwtsecret,
		salt,
		db,
	}
}
