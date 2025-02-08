package service

import (
	infra "IMUbackend/internal/infrastructure"
	"context"

	"goa.design/goa/v3/security"
)

func (s *IMUSrv) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	return infra.JWTAuth(ctx, token, s.jwtsecret)
}