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

// Logger is an interface for logging HTTP request/response metadata
type Logger interface {
	LogRequest(metadata RequestMetadata)
	LogResponse(metadata ResponseMetadata)
}

// Middleware wraps an http.Handler to intercept and log HTTP request/response metadata
type Middleware struct {
	handler http.Handler
	logger  Logger
}

// NewMiddleware creates a new Middleware instance
func NewMiddleware(handler http.Handler, logger Logger) *Middleware {
	return &Middleware{
		handler: handler,
		logger:  logger,
	}
}

// ServeHTTP implements the http.Handler interface
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract request metadata and log it
	// TODO: Wrap response writer to capture response metadata
	// TODO: Call the wrapped handler
	// TODO: Log response metadata
	
	// Placeholder - just pass through for now
	m.handler.ServeHTTP(w, r)
}

// Wrap is a convenience function that wraps an http.Handler with the middleware
func Wrap(handler http.Handler, logger Logger) http.Handler {
	return NewMiddleware(handler, logger)
}
