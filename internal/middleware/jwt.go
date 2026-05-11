package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ctxKey int

const (
	ctxKeyClaims ctxKey = iota
)

// Claims is the JWT payload for both admin and merchant tokens.
type Claims struct {
	Uid    int64    `json:"uid"`
	Role   string   `json:"role"`    // "admin" | "merchant"
	ShopId int64    `json:"shop_id"` // 0 for admin
	Perms  []string `json:"perms"`
	jwt.RegisteredClaims
}

// IssueToken signs a new JWT with the supplied claims.
func IssueToken(uid int64, role string, shopId int64, perms []string, secret string, expire int64) (string, error) {
	now := time.Now()
	c := Claims{
		Uid:    uid,
		Role:   role,
		ShopId: shopId,
		Perms:  perms,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expire) * time.Second)),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tok.SignedString([]byte(secret))
}

// ParseToken validates the bearer token and returns its claims.
func ParseToken(tokenStr, secret string) (*Claims, error) {
	tok, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if c, ok := tok.Claims.(*Claims); ok && tok.Valid {
		return c, nil
	}
	return nil, errors.New("invalid token")
}

// extractBearer pulls the Bearer token from the Authorization header.
func extractBearer(r *http.Request) string {
	h := r.Header.Get("Authorization")
	if h == "" {
		return ""
	}
	if strings.HasPrefix(h, "Bearer ") {
		return strings.TrimPrefix(h, "Bearer ")
	}
	return h
}

// ClaimsFromContext returns the claims stored by the auth middleware.
func ClaimsFromContext(ctx context.Context) (*Claims, bool) {
	c, ok := ctx.Value(ctxKeyClaims).(*Claims)
	return c, ok
}

// withClaims stores claims in the request context.
func withClaims(r *http.Request, c *Claims) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), ctxKeyClaims, c))
}
