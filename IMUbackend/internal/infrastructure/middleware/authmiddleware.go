package middleware

import (
	interceptors "IMUbackend/gen/imubackend"
	"IMUbackend/internal/infrastructure"
	"context"
	"net/http"
	"time"

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
	newCtx, newToken, err := infrastructure.JWTAuth(ctx, token, i.jwtsecret)
	if err != nil {
		return nil, err
	}
	if newToken != "" {
		writer := newCtx.Value("responseWriter").(http.ResponseWriter)
		tokenCookie := http.Cookie{
			Name: "token",
			Value: newToken,
			HttpOnly: true,
			Secure: true,
			SameSite: http.SameSiteStrictMode,
			MaxAge: int(time.Hour * 24),
			Path: "/",
		}
		http.SetCookie(writer, &tokenCookie)
	}
	return next(newCtx, info.RawPayload())
}
