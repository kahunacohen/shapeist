# httpshape

A Go middleware library for the standard `net/http` package that intercepts HTTP requests to log metadata including request/response information and data shape.

## Overview

This library provides a middleware scaffold for logging HTTP request and response metadata. The middleware wraps standard `http.Handler` implementations and can be used to collect information such as:

- Request path, method, and headers
- Remote address and user agent
- Response status code and content length
- Request/response timing

## Installation

```bash
go get github.com/kahunacohen/httpshape
```

## Usage

The middleware is designed to work with any `http.Handler` and requires a logger implementation:

```go
package main

import (
    "net/http"
    "github.com/kahunacohen/httpshape"
)

// Implement the Logger interface
type MyLogger struct{}

func (l *MyLogger) LogRequest(metadata httpshape.RequestMetadata) {
    // Your request logging implementation
}

func (l *MyLogger) LogResponse(metadata httpshape.ResponseMetadata) {
    // Your response logging implementation
}

func main() {
    // Your HTTP handler
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })
    
    // Wrap with middleware
    logger := &MyLogger{}
    wrapped := httpshape.Wrap(handler, logger)
    
    http.ListenAndServe(":8080", wrapped)
}
```

## Architecture

### Middleware Structure

The library provides:
- `RequestMetadata`: Struct containing HTTP request metadata
- `ResponseMetadata`: Struct containing HTTP response metadata
- `Logger`: Interface for implementing custom logging behavior
- `Middleware`: The core middleware type that wraps `http.Handler`
- `Wrap()`: Convenience function for wrapping handlers

### Scaffold Only

**Note**: This is a scaffold implementation. The `ServeHTTP` method currently passes requests through without intercepting metadata. Implementing the actual interception logic is left as an exercise.

## Test Backend

A complete test backend is available in the `/test` subdirectory. This is an in-memory fake healthcare server that demonstrates a RESTful API with CRUD operations for patient entities.

See [test/README.md](test/README.md) for details on running and using the test server.

## License

This is a reference implementation for educational purposes.
