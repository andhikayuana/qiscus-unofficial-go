package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	appCode   = "test-app-code"
	secretKey = "test-secret-key"
)

func TestNewSDK(t *testing.T) {
	c := NewSDK(appCode, secretKey)
	assert.Equal(t, c.AppCode(), appCode)
	assert.Equal(t, c.SecretKey(), secretKey)
}
