package qiscus

import (
	"net/http"
	"time"
)

// LibraryVersion : qiscus go library version
const LibraryVersion = "v1.0.3"

var (
	// DefaultHttpTimeout default timeout for go HTTP HttpClient
	DefaultHttpTimeout = 80 * time.Second

	// DefaultGoHttpClient default go HTTP Client for Qiscus HttpClient API
	DefaultHttpClient = &http.Client{Timeout: DefaultHttpTimeout}

	// DefaultOutboundLog default HTTP outbound log
	DefaultHttpOutboundLog = false
)
