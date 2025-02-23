package middleware

import (
	interceptors "IMUbackend/gen/imubackend"
	"IMUbackend/internal/infrastructure"
	"context"

	goa "goa.design/goa/v3/pkg"
)

type Interceptor struct {
	jwtsecret string
}

func NewInterceptor(jwtsecret string) *Interceptor {
	return &Interceptor{
		jwtsecret: jwtsecret,
	}
}

func (i *Interceptor) JWTAuth(ctx context.Context, info *interceptors.JWTAuthInfo, next goa.Endpoint) (any, error) {
	token := info.Payload().Token()
	newCtx, _, err := infrastructure.JWTAuth(ctx, token, i.jwtsecret)
	if err != nil {
		return nil, err
	}
	return next(newCtx, info.RawPayload())
}
