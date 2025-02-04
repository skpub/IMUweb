package service

import (
	pb "IMUbackend/gen/imubackend"
	repo "IMUbackend/internal/repository"
)

type IMUSrv struct {
	article   repo.IArticleRepository
	user repo.IStudentRepository
	tx   repo.TxManager
}

func NewIMUSrv(article repo.IArticleRepository, user repo.IStudentRepository, txManager repo.TxManager) pb.Service {
	return &IMUSrv{
		article,
		user,
		txManager,
	}
}
