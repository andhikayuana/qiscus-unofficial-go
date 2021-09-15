package multichannel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	appCode    = "test-app-code"
	adminToken = "test-admin-token"
	adminEmail = "test@mail.com"
)

func TestNewMultichannel(t *testing.T) {
	c := NewMultichannel(appCode, adminToken, adminEmail)

	assert.Equal(t, c.AppCode(), appCode)
	assert.Equal(t, c.AdminToken(), adminToken)
	assert.Equal(t, c.AdminEmail(), adminEmail)
}
