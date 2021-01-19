package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/ivyoverflow/pub-sub/book/internal/logger"
)

// Middleware ...
type Middleware struct {
	ctx context.Context
	log *logger.Logger
}

// NewMiddleware ...
func NewMiddleware(ctx context.Context, log *logger.Logger) *Middleware {
	return &Middleware{ctx, log}
}

// AbortWithContext ...
func (mw *Middleware) AbortWithContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		select {
		case <-ctx.Done():
			mw.log.Debug("Aborting connection...")
			rw.WriteHeader(http.StatusServiceUnavailable)
			cancel()

			return
		default:
			next.ServeHTTP(rw, r)
		}

	})
}
