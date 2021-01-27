package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/ivyoverflow/pub-sub/book/internal/logger"
)

// Middleware contains all middleware methods.
type Middleware struct {
	ctx context.Context
	log *logger.Logger
}

// NewMiddleware returns a new configured Middleware object.
func NewMiddleware(ctx context.Context, log *logger.Logger) *Middleware {
	return &Middleware{ctx, log}
}

// AbortWithContext shutdowns the connection with timeout.
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
