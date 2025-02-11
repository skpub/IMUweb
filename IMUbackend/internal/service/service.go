package service

import (
	pb "IMUbackend/gen/imubackend"
	repo "IMUbackend/internal/repository"
	infra "IMUbackend/internal/infrastructure"
)


type IMUSrv struct {
	article   repo.IArticleRepository
	user      repo.IStudentRepository
	jwtsecret string
	salt      string
	db        infra.IDBTX
}


func NewIMUSrv(article repo.IArticleRepository, user repo.IStudentRepository, jwtsecret string, salt string, db infra.IDBTX) pb.Service {
	return &IMUSrv{
		article,
		user,
		jwtsecret,
		salt,
		db,
	}
}
