package infrastructure

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestJWTAuth_CorruptedToken(t *testing.T) {
	secret := "secret"
	tokenString := "CORRUPTED"
	ctx := context.Background()
	_, _, err := JWTAuth(ctx, tokenString, secret)
	assert.Error(t, err)
}

func TestJWTAuth_ExpiredToken(t *testing.T) {
	secret := "secret"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwic3R1ZGVudF9pZCI6InRlc3QiLCJleHAiOjIxMX0.GqzAlQDGaqkkbG20BpjQDOhiiaNG5_FbaPFckFUvyl8"
	ctx := context.Background()
	_, _, err := JWTAuth(ctx, tokenString, secret)
	assert.Error(t, err)
}

func TestJWTAuth_ValidToken(t *testing.T) {
	secret := "secret"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwic3R1ZGVudF9pZCI6InRlc3QiLCJleHAiOjk5OTk5OTk5OTk5OX0.ALvyDIegW3hLXGQwfMpz7Pwadoyo5QZd-eT7nHvwCQo"
	ctx := context.Background()
	_, _, err := JWTAuth(ctx, tokenString, secret)
	assert.NoError(t, err)
}
