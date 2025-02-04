package service

import (
	pb "IMUbackend/gen/imubackend"
	"context"
	"IMUbackend/db"
)

// returns (token, err)
func (s *IMUSrv) Login(ctx context.Context, attribute *pb.Login2) (string, error) {
	err := s.user.Login(ctx, nil, *attribute.StudentName, *attribute.Password)
	if err != nil {
		return "", err
	}
	return "", nil	
}

func (s *IMUSrv) CreateStudent(ctx context.Context, attribute *pb.Signup) error {
	tx, err := s.txManager.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	err = s.user.Create(ctx, tx, db.Student{
		ID: 		*attribute.StudentID,
		Name: 		*attribute.StudentName,
		Password: 	*attribute.Password,
		Email: 		*attribute.Email,
	})
	if err != nil {
		return err
	}
	return nil
}
