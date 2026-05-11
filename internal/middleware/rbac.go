package middleware

import (
	"net/http"

	"mall-common/configcenter"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// RoleMiddleware enforces JWT auth and a required role.
type RoleMiddleware struct {
	secret       *configcenter.HotConfig[string]
	requiredRole string
}

func NewRoleMiddleware(secret *configcenter.HotConfig[string], requiredRole string) *RoleMiddleware {
	return &RoleMiddleware{secret: secret, requiredRole: requiredRole}
}

func (m *RoleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tok := extractBearer(r)
		if tok == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]any{"code": 1003, "msg": "missing token"})
			return
		}
		claims, err := ParseToken(tok, m.secret.Get())
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, map[string]any{"code": 1003, "msg": "invalid token"})
			return
		}
		if claims.Role != m.requiredRole {
			httpx.WriteJson(w, http.StatusForbidden, map[string]any{"code": 1003, "msg": "forbidden"})
			return
		}
		next(w, withClaims(r, claims))
	}
}
