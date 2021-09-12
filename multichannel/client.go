package multichannel

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/andhikayuana/qiscus-unofficial-go"
)

// Base Url the library uses to contact multichannel. Use SetAPIBase() to override
const APIBase = "https://multichannel.qiscus.com"

// Multichannel defines the supported subset of the Multichannel API.
type Multichannel interface {
	APIBase() string
	AppCode() string
	AdminToken() string
	AdminEmail() string
	SetAPIBase(address string)

	// Room
	GetRoomTags(roomID string) (*RoomTagsResponse, *qiscus.Error)
	CreateRoomTag(roomID, tag string) (*CreateRoomTagResponse, *qiscus.Error)
	CreateAdditionalInfoRoomWithReplace(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error)
	GetAdditionalInfoRoom(roomID string) (*GetAdditionalInfoRoomResponse, *qiscus.Error)
	CreateAdditionalInfoRoom(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error)
	MarkAsResolved(req *MarkAsResolvedReq) (*MarkAsResolvedResponse, *qiscus.Error)
	GetRoomByRoomID(roomID string) (*GetRoomByRoomIDResponse, *qiscus.Error)

	// BOT
	SendMessageTextByBot(roomID, message string) *qiscus.Error
	SetToogleBotInRoom(roomID string, isActive bool) (*SetToogleBotInRoomResponse, *qiscus.Error)

	// Agent & Channel
	GetAllAgents(req *GetAllAgentsReq) (*GetAllAgentsResponse, *qiscus.Error)
	AssignAgent(req *AssignAgentReq) (*AssignAgentResponse, *qiscus.Error)
	GetAgentsByDivision(req *GetAgentsByDivisionReq) (*GetAgentsByDivisionResponse, *qiscus.Error)
	GetAllDivision(req *GetAllDivisionReq) (*GetAllDivisionResponse, *qiscus.Error)
	GetAllChannels() (*GetAllChannelsResponse, *qiscus.Error)
}

// // MultichannelImpl bundles data needed by a large number of methods in order to interact with the Multichannel API.
type MultichannelImpl struct {
	apiBase    string
	appCode    string
	adminToken string
	adminEmail string
}

// NewMultichannel creates a new client instance.
func NewMultichannel(appCode, adminToken, adminEmail string) Multichannel {
	return &MultichannelImpl{
		apiBase:    APIBase,
		appCode:    appCode,
		adminToken: adminToken,
		adminEmail: adminEmail,
	}
}

// NewMultichannelFromEnv returns a new Multichannel client using the environment variables
// QISMO_APP_CODE, QISMO_ADMIN_TOKEN, QISMO_ADMIN_EMAIL and QISMO_BASE_URL
func NewMultichannelFromEnv() (Multichannel, error) {
	appCode := os.Getenv("QISMO_APP_CODE")
	if appCode == "" {
		return nil, errors.New("required environment variable QISMO_APP_CODE not defined")
	}

	adminToken := os.Getenv("QISMO_ADMIN_TOKEN")
	if adminToken == "" {
		return nil, errors.New("required environment variable QISMO_ADMIN_TOKEN not defined")
	}

	adminEmail := os.Getenv("QISMO_ADMIN_EMAIL")
	if adminEmail == "" {
		return nil, errors.New("required environment variable QISMO_ADMIN_EMAIL not defined")
	}

	m := NewMultichannel(appCode, adminToken, adminEmail)

	url := os.Getenv("QISMO_BASE_URL")
	if url != "" {
		m.SetAPIBase(url)
	}

	return m, nil

}

func NewMultichannelFromCredential(email, password string) (Multichannel, error) {
	resp := &LoginAdminResponse{}
	url := fmt.Sprintf("%s/api/v1/auth", APIBase)

	req := &LoginAdminReq{Email: email, Password: password}
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	if err := r.DoRequest(); err != nil {
		return nil, fmt.Errorf("initiate client for multichannel failed. %s", err.Message)
	}

	m := NewMultichannel(resp.Data.User.App.AppCode, resp.Data.LongLivedToken, resp.Data.User.SdkEmail)
	return m, nil

}

// APIBase returns the API Base URL configured for this client
func (m *MultichannelImpl) APIBase() string {
	return m.apiBase
}

// AppCode returns the App ID configured for this client
func (m *MultichannelImpl) AppCode() string {
	return m.appCode
}

// AppID returns the Admin Token configured for this client
func (m *MultichannelImpl) AdminToken() string {
	return m.adminToken
}

// AppID returns the Admin Email configured for this client
func (m *MultichannelImpl) AdminEmail() string {
	return m.adminEmail
}

// SetAPIBase updates the API Base URL for this client.
// Set a custom base API: m.SetAPIBase("https://multichannel-test.qiscus.com")
func (m *MultichannelImpl) SetAPIBase(address string) {
	m.apiBase = address
}
