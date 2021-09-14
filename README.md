# Unofficial Qiscus API for Golang
[![Go Reference](https://pkg.go.dev/badge/github.com/andhikayuana/qiscus-unofficial-go.svg)](https://pkg.go.dev/github.com/andhikayuana/qiscus-unofficial-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/andhikayuana/qiscus-unofficial-go)](https://goreportcard.com/report/github.com/andhikayuana/qiscus-unofficial-go)

This library is the abstraction of [Qiscus](https://www.qiscus.com) SDK & Multichannel API for access from applications written with Go.

## 1. Installation
Install qiscus-unofficial-go with:

```sh
go get -u github.com/andhikayuana/qiscus-unofficial-go
```

Then, import it using:

```go
import (
    "github.com/andhikayuana/qiscus-unofficial-go"
    "github.com/andhikayuana/qiscus-unofficial-go/$product$"
)
```
with `$product$` is the product of Qiscus such as `sdk` and `multichannel`.

## 2. Usage
```go
package main

import (
	"fmt"

	"github.com/andhikayuana/qiscus-unofficial-go/multichannel"
	"github.com/andhikayuana/qiscus-unofficial-go/sdk"
)

func main() {
	// Initiate client for Multichannel.
	multichannelClient := multichannel.NewMultichannel("app-code", "admin-token", "admin-email")

	// Initiate client for Multichannel using creadential email and password admin.
	multichannelClient, err := multichannel.NewMultichannelFromCredential("example@mail.com", "password")
	if err != nil {
		panic(err)
	}

	// Initiate client for Multichannel using environment variable.
	// QISMO_APP_CODE, QISMO_ADMIN_TOKEN, QISMO_ADMIN_EMAIL and QISMO_BASE_URL --optional
	multichannelClient, err := multichannel.NewMultichannelFromEnv()
	if err != nil {
		panic(err)
	}

	// Sample Multichannel method.
	resp, _ := multichannelClient.GetRoomTags("12345678")
	fmt.Println(resp)


	// Initiate client for SDK.
	sdkClient := sdk.NewSDK("app-code", "secret-key")
	
	// Initiate client for SDK using environment variable.
	// QISCUS_SDK_APP_CODE, QISCUS_SDK_SECRET_KEY and QISCUS_SDK_BASE_URL --optional
	sdkClient, err := sdk.NewSDKFromEnv()
	if err != nil {
		panic(err)
	}

	// Sample SDK method.
	resp, _ := sdkClient.LoginOrRegister(&sdk.LoginOrRegisterReq{
		UserID:   "guest@qiscus.com",
		Password: "12345678",
		Username: "User Demo",
	})
	fmt.Println(resp)

}

```

## 3. Advance Usage
### 3.1. Override Base API URL
```go
sdkClient,_ := sdk.NewSDKFromEnv()
// Default SDK base is https://api.qiscus.com, you can use SetAPIBase() to override.
sdkClient.SetAPIBase("https://api2.qiscus.com")

multichannelClient, _ := multichannel.NewMultichannelFromEnv()
// Default Multichannel base is https://multichannel.qiscus.com, you can use SetAPIBase() to override.
multichannelClient.SetAPIBase("https://multichannel2.qiscus.com")
```

### 3.2. Override HTTP Client timeout
By default, timeout value for HTTP Client 80 seconds. But you can override the HTTP client default config from global variable `qiscus.DefaultGoHttpClient`:
```go
t := 100 * time.Second
qiscus.DefaultGoHttpClient = &http.Client{
	Timeout: t,
}
```

### 3.3. Log Configuration
By default, the log level will use `LogError` level. You have option to change the default log level configuration with global variable `qiscus.DefaultLoggerLevel`:
```go
qiscus.DefaultLoggerLevel = &qiscus.LoggerImpl{LogLevel: qiscus.LogInfo}

// Details Log Level
// NoLogging    : sets a logger to not show the messages
// LogError     : sets a logger to show error messages only.
// LogInfo      : sets a logger to show information messages
// LogDebug     : sets a logger to show informational messages for debugging
```

## 4. Error Handling
Several functions in the product allow to throw an error, below is an qiscus error object you can use:
```go
_, err := multichannelClient.GetRoomTags("12345678")
if err != nil {
	message := err.GetMessage()               // general message error
	statusCode := err.GetStatusCode()         // HTTP status code e.g: 400, 401, etc.
	rawApiResponse := err.GetRawApiResponse() // raw Go HTTP response object
	rawError := err.GetRawError()             // raw Go err object
}
```
