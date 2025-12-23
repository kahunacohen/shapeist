// Package httpshape provides structures for HTTP request and response metadata.
package httpshape

import (
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
