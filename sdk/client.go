package sdk

import (
	"errors"
	"os"

	"github.com/syahidfrd/go-qiscus"
)

// Base Url the library uses to contact multichannel. Use SetAPIBase() to override
const APIBase = "https://api.qiscus.com"

// SDK defines the supported subset of the SDK API.
type SDK interface {
	APIBase() string
	AppCode() string
	SecretKey() string
	SetAPIBase(address string)

	LoginOrRegister(req *LoginOrRegisterReq) (*LoginOrRegisterResponse, *qiscus.Error)
	GetUserProfile(userID string) (*GetUserProfileResponse, *qiscus.Error)
	GetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error)
	ResetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error)
	CreateRoom(req *CreateRoomReq) (*CreateRoomResponse, *qiscus.Error)
	GetOrCreateRoomWithTarget(req *GetOrCreateRoomWithTargetReq) (*CreateRoomResponse, *qiscus.Error)
	GetRoomsInfo(roomIDs []string) (*GetRoomsInfoResponse, *qiscus.Error)
	UpdateRoom(req *UpdateRoomReq) (*UpdateRoomResponse, *qiscus.Error)
	GetRoomParticipants(roomID string, page, limit int) (*GetRoomParticipantsResponse, *qiscus.Error)
	AddRoomParticipants(req *AddRoomParticipantsReq) (*AddRoomParticipantsResponse, *qiscus.Error)
	RemoveRoomParticipants(req *RemoveRoomParticipantsReq) (*RemoveRoomParticipantsResponse, *qiscus.Error)
	GetUserRooms(userID string, page, limit int) (*GetUserRoomsResponse, *qiscus.Error)
	PostComment(req *PostCommentReq) (*PostCommentResponse, *qiscus.Error)
}

// SDKImpl bundles data needed by a large number of methods in order to interact with the SDK API.
type SDKImpl struct {
	apiBase   string
	appCode   string
	secretKey string
}

// NewSDK creates a new client instance
func NewSDK(appCode, secretKey string) SDK {
	return &SDKImpl{
		apiBase:   APIBase,
		appCode:   appCode,
		secretKey: secretKey,
	}
}

// NewSDKFromEnv returns a new SDK client using the environment variables
// QISCUS_SDK_APP_CODE, QISCUS_SDK_SECRET_KEY and QISCUS_SDK_BASE_URL
func NewSDKFromEnv() (SDK, error) {
	appCode := os.Getenv("QISCUS_SDK_APP_CODE")
	if appCode == "" {
		return nil, errors.New("required environment variable QISCUS_SDK_APP_CODE not defined")
	}

	secretKey := os.Getenv("QISCUS_SDK_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("required environment variable QISCUS_SDK_SECRET_KEY not defined")
	}

	s := NewSDK(appCode, secretKey)

	url := os.Getenv("QISCUS_SDK_BASE_URL")
	if url != "" {
		s.SetAPIBase(url)
	}

	return s, nil
}

// APIBase returns the API Base URL configured for this client
func (s *SDKImpl) APIBase() string {
	return s.apiBase
}

// SecretKey returns the API Base URL configured for this client
func (s *SDKImpl) SecretKey() string {
	return s.secretKey
}

// AppCode returns the App ID configured for this client
func (s *SDKImpl) AppCode() string {
	return s.appCode
}

// SetAPIBase updates the API Base URL for this client.
// Set a custom base API: m.SetAPIBase("https://api3.qiscus.com")
func (s *SDKImpl) SetAPIBase(address string) {
	s.apiBase = address
}
