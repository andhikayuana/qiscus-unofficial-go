package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andhikayuana/qiscus-unofficial-go"
)

// LoginOrRegister Login or register
func (s *SDKImpl) LoginOrRegister(req *LoginOrRegisterReq) (*LoginOrRegisterResponse, *qiscus.Error) {
	resp := &LoginOrRegisterResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/login_or_register", s.APIBase())

	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetUserProfile Get user profile by user ID
func (s *SDKImpl) GetUserProfile(userID string) (*GetUserProfileResponse, *qiscus.Error) {
	resp := &GetUserProfileResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/user_profile", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	r.AddParameter("user_id", userID)
	err := r.DoRequest()

	return resp, err
}

// GetUserToken Get user profile by user ID
func (s *SDKImpl) GetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error) {
	resp := &GetUserTokenResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_user_token", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	r.AddParameter("user_id", userID)
	err := r.DoRequest()

	return resp, err
}

// ReserUserToken Reset user token by user ID
func (s *SDKImpl) ResetUserToken(userID string) (*GetUserTokenResponse, *qiscus.Error) {
	resp := &GetUserTokenResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/reset_user_token", s.APIBase())

	req := &ResetUserTokenReq{UserID: userID}
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	err := r.DoRequest()

	return resp, err
}

// CreateRoom Create new room
func (s *SDKImpl) CreateRoom(req *CreateRoomReq) (*CreateRoomResponse, *qiscus.Error) {
	resp := &CreateRoomResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/create_room", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetOrCreateRoomWithTarget Get or create new room with target
func (s *SDKImpl) GetOrCreateRoomWithTarget(req *GetOrCreateRoomWithTargetReq) (*CreateRoomResponse, *qiscus.Error) {
	resp := &CreateRoomResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_or_create_room_with_target", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetRoomsInfo Get rooms info by room IDs
func (s *SDKImpl) GetRoomsInfo(roomIDs []string) (*GetRoomsInfoResponse, *qiscus.Error) {
	resp := &GetRoomsInfoResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_rooms_info", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())

	for _, roomID := range roomIDs {
		r.AddParameter("room_ids[]", roomID)
	}

	err := r.DoRequest()

	return resp, err
}

// UpdateRoom Update room
func (s *SDKImpl) UpdateRoom(req *UpdateRoomReq) (*UpdateRoomResponse, *qiscus.Error) {
	resp := &UpdateRoomResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/update_room", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetRoomParticipants: Represent Get room participant
func (s *SDKImpl) GetRoomParticipants(roomID string, page, limit int) (*GetRoomParticipantsResponse, *qiscus.Error) {
	resp := &GetRoomParticipantsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_room_participants", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	r.AddParameter("room_id", roomID)
	r.AddParameter("page", strconv.Itoa(page))
	r.AddParameter("limit", strconv.Itoa(limit))
	err := r.DoRequest()

	return resp, err
}

// AddRoomParticipants Add room participants
func (s *SDKImpl) AddRoomParticipants(req *AddRoomParticipantsReq) (*AddRoomParticipantsResponse, *qiscus.Error) {
	resp := &AddRoomParticipantsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/add_room_participants", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	err := r.DoRequest()

	return resp, err
}

// RemoveRoomParticipants Remove room participants
func (s *SDKImpl) RemoveRoomParticipants(req *RemoveRoomParticipantsReq) (*RemoveRoomParticipantsResponse, *qiscus.Error) {
	resp := &RemoveRoomParticipantsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/remove_room_participants", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	err := r.DoRequest()

	return resp, err
}

// GetUserRooms Get user rooms
func (s *SDKImpl) GetUserRooms(userID string, page, limit int) (*GetUserRoomsResponse, *qiscus.Error) {
	resp := &GetUserRoomsResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/get_user_rooms", s.APIBase())

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	r.AddParameter("user_id", userID)
	r.AddParameter("page", strconv.Itoa(page))
	r.AddParameter("limit", strconv.Itoa(limit))
	err := r.DoRequest()

	return resp, err

}

// PostComment Post comment
func (s *SDKImpl) PostComment(req *PostCommentReq) (*PostCommentResponse, *qiscus.Error) {
	resp := &PostCommentResponse{}
	url := fmt.Sprintf("%s/api/v2.1/rest/post_comment", s.APIBase())
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("QISCUS_SDK_APP_ID", s.AppCode())
	r.AddHeader("QISCUS_SDK_SECRET", s.SecretKey())
	err := r.DoRequest()

	return resp, err
}
