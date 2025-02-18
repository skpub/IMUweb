package service

import (
	"IMUbackend/db"
	pb "IMUbackend/gen/imubackend"
	"IMUbackend/internal/infrastructure"
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
	err = s.user.Login(ctx, *attribute.StudentID, pwhashstr)
	if err != nil {
		return "", err
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"student_id": *attribute.StudentID,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.jwtsecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *IMUSrv) RefreshToken(ctx context.Context, token *pb.RefreshTokenPayload) (*pb.RefreshTokenResult, error) {
	tokenStr := token.Token
	_, studentID, err := infrastructure.JWTAuth(ctx, *tokenStr, s.jwtsecret)
	if err != nil {
		return nil, err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"student_id": studentID,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	refreshTokenStr, err := refreshToken.SignedString([]byte(s.jwtsecret))
	if err != nil {
		return nil, err
	}
	result := &pb.RefreshTokenResult{
		Token: &refreshTokenStr,
	}
	return result, nil
}

func (s *IMUSrv) Signup(ctx context.Context, attribute *pb.SignupPayload) (string, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			s.db.Rollback(tx)
		} else {
			s.db.Commit(tx)
		}
	}()

	pwhashstr, err := genhash(*attribute.Password, s.salt)
	if err != nil {
		return "", err
	}

	id, err := s.user.Create(ctx, db.CreateStudentParams {
		ID: 	 *attribute.StudentID,
		Name:    *attribute.Name,
		Email:   *attribute.Email,
		Password: pwhashstr,
	})
	return id, err
}
