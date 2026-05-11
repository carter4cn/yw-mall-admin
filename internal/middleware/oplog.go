package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// OpLogMiddleware records admin/merchant write operations via logx.
// Using structured logs lets ops aggregate via existing pipelines without
// taking a cross-database write dependency in the gateway.
type OpLogMiddleware struct{}

func NewOpLogMiddleware() *OpLogMiddleware { return &OpLogMiddleware{} }

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func (m *OpLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := strings.ToUpper(r.Method)
		if method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions {
			next(w, r)
			return
		}
		var snippet []byte
		if r.Body != nil {
			b, _ := io.ReadAll(r.Body)
			snippet = b
			r.Body = io.NopCloser(bytes.NewBuffer(b))
		}
		if len(snippet) > 4096 {
			snippet = snippet[:4096]
		}
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		start := time.Now()
		next(rec, r)
		var actorId int64
		actorRole := ""
		if c, ok := ClaimsFromContext(r.Context()); ok {
			actorId = c.Uid
			actorRole = c.Role
		}
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}
		logx.WithContext(r.Context()).Infof(
			"[oplog] actor=%d role=%s %s %s status=%d ip=%s elapsed=%s body=%s",
			actorId, actorRole, method, r.URL.Path, rec.status, ip, time.Since(start), string(snippet))
	}
}
