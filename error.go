package qiscus

type Error struct {
	Message        string
	StatusCode     int
	RawError       error
	RawApiResponse *APIResponse
}

// GetMessage this get general message error when call api
func (e *Error) GetMessage() string {
	return e.Message
}

// GetStatusCode this get api response status code coming from midtrans backend
func (e *Error) GetStatusCode() int {
	return e.StatusCode
}

// GetRawApiResponse this get api raw response from midtrans backend
func (e *Error) GetRawApiResponse() *APIResponse {
	return e.RawApiResponse
}

// GetRawError GetRawApiResponse this get api raw response from midtrans backend
func (e *Error) GetRawError() error {
	return e.RawError
}
