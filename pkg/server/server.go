package server

import (
	"fmt"
	"net/http"
	"time"
)

type serverOption struct {
	writeTimeout time.Duration
	readTimeout  time.Duration
}

// ServerOption provides possibility of configuring the server by the caller
// can be extended to have other attributes
type ServerOption func(*serverOption)

func WithWriteTimeout(d time.Duration) ServerOption {
	return func(opt *serverOption) {
		opt.writeTimeout = d
	}
}

func WithReadTimeout(d time.Duration) ServerOption {
	return func(opt *serverOption) {
		opt.readTimeout = d
	}
}

// NewServer runs the server
func NewServer(addr int, router http.Handler, opts ...ServerOption) *http.Server {
	options := &serverOption{
		writeTimeout: 15 * time.Second,
		readTimeout:  15 * time.Second,
	}

	for _, opt := range opts {
		opt(options)
	}

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", addr),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: options.writeTimeout,
		ReadTimeout:  options.readTimeout,
	}
}
