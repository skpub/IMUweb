package service

import (
	"IMUbackend/db"
	pb "IMUbackend/gen/imubackend"
	"context"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/scrypt"
)

func genhash(password string, salt string) (string, error) {
	pwhash, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	pwhashStr := base64.StdEncoding.EncodeToString(pwhash)
	return pwhashStr, nil
}

// returns (token, err)
func (s *IMUSrv) Login(ctx context.Context, attribute *pb.Login2) (string, error) {
	pwhashstr, err := genhash(*attribute.Password, s.salt)
	if err != nil {
		return "", err
	}

	// DB query
	err = s.user.Login(ctx, *attribute.StudentName, pwhashstr)
	if err != nil {
		return "", err
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"student_id": *attribute.StudentName,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.jwtsecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *IMUSrv) CreateStudent(ctx context.Context, attribute *pb.Signup) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	pwhashstr, err := genhash(*attribute.Password, s.salt)
	if err != nil {
		return err
	}

	return s.user.Create(ctx, db.Student{
		ID:       *attribute.StudentID,
		Name:     *attribute.StudentName,
		Password: pwhashstr,
		Email:    *attribute.Email,
	})
}
