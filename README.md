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

	"github.com/andhikayuana/qiscus-unofficial-go"
	"github.com/andhikayuana/qiscus-unofficial-go/multichannel"
	"github.com/andhikayuana/qiscus-unofficial-go/sdk"
)

func main() {
	qiscus.DefaultHttpOutboundLog = true

	// Initiate client for Multichannel.
	multichannelClient := multichannel.NewMultichannel("qiscus-app-id", "qiscus-secret-key")

	// Initiate client for Multichannel using creadential email and password admin.
	multichannelClient, err := multichannel.NewMultichannelFromCredential("example@mail.com", "12345678")
	if err != nil {
		panic(err)
	}

	// Initiate client for Multichannel using environment variable.
	// QISCUS_APP_ID, QISCUS_SECRET_KEY and MULTICHANNEL_API_BASE --optional
	multichannelClient, err := multichannel.NewMultichannelFromEnv()
	if err != nil {
		panic(err)
	}

	// Sample Multichannel method.
	resp, _ := multichannelClient.CreateRoomTag(&multichannel.CreateRoomTagReq{
		RoomID: "12345678",
		Tag:    "test",
	})
	fmt.Println(resp)


	// Initiate client for SDK.
	sdkClient := sdk.NewSDK("qiscus-app-id", "qiscus-secret-key")
	
	// Initiate client for SDK using environment variable.
	// QISCUS_APP_ID, QISCUS_SECRET_KEY and QISCUS_API_BASE --optional
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
By default, timeout value for HTTP Client 80 seconds. But you can override the HTTP client default config from global variable `qiscus.DefaultHttpClient`:
```go
t := 100 * time.Second
qiscus.DefaultHttpClient = &http.Client{
	Timeout: t,
}
```

### 3.3. HTTP Outbound Log Configuration
By default, the outbound log is `false`. You have option to change the default outbound log configuration with global variable `qiscus.DefaultHttpOutboundLog`:
```go
qiscus.DefaultHttpOutboundLog = true

// Details HTTP Outbound Log
{
  "level": "info",
  "method": "POST",
  "url": "https://multichannel.qiscus.com/api/v1/room_tag/create",
  "body": "{\"room_id\":\"12345678\",\"tag\":\"test\"}",
  "status": 200,
  "response": "{\"data\":{\"id\":1,\"name\":\"test\"}}",
  "latency": 774.9559,
  "time": "2021-09-20T14:32:24+07:00",
  "message": "OUTBOUND LOG"
}
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
