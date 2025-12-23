// Package shapeist provides structures for HTTP request and response metadata.
package shapeist

import (
	"fmt"
	"math/rand"
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

type Middleware struct {
	sampleRate float64
}

func NewMiddleware(sampleRate float64) *Middleware {
	return &Middleware{
		sampleRate: sampleRate,
	}
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func shouldSample(rate float64) bool {
	return rng.Float64() < rate
}

func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if shouldSample(m.sampleRate) {
			fmt.Println("take a sample of", r.Method, r.URL.Path)
		}
		next.ServeHTTP(w, r)
	})
}
