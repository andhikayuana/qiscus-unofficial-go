package qiscus

import (
	"net/http"
	"time"
)

// LibraryVersion : qiscus go library version
const LibraryVersion = "v1.0.0"

var (
	//DefaultHttpTimeout default timeout for go HTTP HttpClient
	DefaultHttpTimeout = 80 * time.Second

	//DefaultGoHttpClient default Go HTTP Client for Qiscus HttpClient API
	DefaultGoHttpClient = &http.Client{Timeout: DefaultHttpTimeout}
)
