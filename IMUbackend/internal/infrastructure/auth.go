package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 2nd return vaulue: new token or ""
func JWTAuth(ctx context.Context, tokenString string, secret string) (context.Context, string,  error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		rawExp, ok := claims["exp"].(float64)
		if !ok {
			return nil, "", fmt.Errorf("invalid token. \"exp\" field is of invalid type")
		}
		exp := int64(rawExp)
		if exp < time.Now().Unix() {
			return nil, "", fmt.Errorf("token is expired")

		} else if exp - time.Now().Unix() < 60 * 60 * 24 {
			// refresh token
			student_id := claims["student_id"].(string)
			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"student_id": student_id,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
			})
			tokenStr, err := newToken.SignedString([]byte(secret))
			if err != nil {
				return nil, "", err
			}
			return context.WithValue(ctx, "studentId", student_id), tokenStr, nil
		} else {
			student_id := claims["student_id"].(string)
			return context.WithValue(ctx, "studentId", student_id), "", nil
		}
	}
	return nil, "", fmt.Errorf("invalid token")
}
