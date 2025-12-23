// Package httpshape provides structures for HTTP request and response metadata.
package httpshape

import (
	"net/http"
	"time"
)

// RequestMetadata holds metadata about an HTTP request
type RequestMetadata struct {
	Method      string
	Path        string
	RemoteAddr  string
	UserAgent   string
	ContentType string
	Timestamp   time.Time
}

// ResponseMetadata holds metadata about an HTTP response
type ResponseMetadata struct {
	StatusCode    int
	ContentLength int64
	Duration      time.Duration
}

func shouldSample(rate float64) bool {
	// Implement sampling logic based on the rate
	return true
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
