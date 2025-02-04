package service

import (
	pb "IMUbackend/gen/imubackend"
	repo "IMUbackend/internal/repository"
)

type IMUSrv struct {
	txManager	repo.TxManager
	md 			repo.IMarkdownRepository
	user 		repo.IUserRepository
}

func NewIMUSrv(md repo.IMarkdownRepository, user repo.IUserRepository, txManager repo.TxManager) pb.Service {
	return &IMUSrv{
		txManager,
		md,
		user,
	}
}


