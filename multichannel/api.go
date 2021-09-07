package multichannel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syahidfrd/go-qiscus"
)

// GetRoomTags get room tags by room ID
func (m *MultichannelImpl) GetRoomTags(roomID string) (*RoomTagsResponse, *qiscus.Error) {
	resp := &RoomTagsResponse{}
	url := fmt.Sprintf("%s/api/v1/room_tag/%s", m.APIBase(), roomID)

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Authorization", m.AdminToken())
	err := r.DoRequest()

	return resp, err
}

// CreateRoomTag create room tag
func (m *MultichannelImpl) CreateRoomTag(roomID, tag string) (*CreateRoomTagResponse, *qiscus.Error) {
	resp := &CreateRoomTagResponse{}
	url := fmt.Sprintf("%s/api/v1/room_tag/create", m.APIBase())

	req := CreateRoomTagReq{RoomID: roomID, Tag: tag}
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Authorization", m.AdminToken())
	err := r.DoRequest()

	return resp, err
}

// CreateAdditionalInfoRoomWithReplace create additional info room with replace exisiting data
func (m *MultichannelImpl) CreateAdditionalInfoRoomWithReplace(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error) {
	resp := &CreateAdditionalInfoRoomResponse{}
	url := fmt.Sprintf("%s/api/v1/qiscus/room/%s/user_info", m.APIBase(), roomID)

	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Authorization", m.AdminToken())
	err := r.DoRequest()

	return resp, err
}

// GetAdditionalInfoRoom get additional info room by room ID
func (m *MultichannelImpl) GetAdditionalInfoRoom(roomID string) (*GetAdditionalInfoRoomResponse, *qiscus.Error) {
	resp := &GetAdditionalInfoRoomResponse{}
	url := fmt.Sprintf("%s/api/v1/qiscus/room/%s/user_info", m.APIBase(), roomID)

	r := qiscus.NewHttpRequest(http.MethodGet, url, nil, resp)
	r.AddHeader("Authorization", m.AdminToken())
	err := r.DoRequest()

	return resp, err
}

// CreateAdditionalInfoRoom create additional info room without replace exisiting data
func (m *MultichannelImpl) CreateAdditionalInfoRoom(roomID string, req *CreateAdditionalInfoRoomReq) (*CreateAdditionalInfoRoomResponse, *qiscus.Error) {
	resp := &CreateAdditionalInfoRoomResponse{}

	res, e := m.GetAdditionalInfoRoom(roomID)
	if e != nil {
		return resp, e
	}

	// Merge existing additional info data (if available).
	// We must do the get and set method when adding data to the additional info room,
	// so as not to replace the existing data.
	existingAdditionalInfoData := []UserProperty{}
	userProperties := res.Data.Extras.UserProperties
	if len(userProperties) > 0 {
		for _, userProp := range userProperties {
			existingAdditionalInfoData = append(existingAdditionalInfoData, UserProperty{
				Key:   userProp.Key,
				Value: userProp.Value,
			})
		}
	}
	req.UserProperties = append(req.UserProperties, existingAdditionalInfoData...)

	url := fmt.Sprintf("%s/api/v1/qiscus/room/%s/user_info", m.APIBase(), roomID)
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Authorization", m.AdminToken())
	err := r.DoRequest()

	return resp, err
}

// SendMessageTextByBot send message text by bot
func (m *MultichannelImpl) SendMessageTextByBot(roomID, message string) *qiscus.Error {
	url := fmt.Sprintf("%s/%s/bot", m.APIBase(), m.AppCode())

	req := SendMessageTextByBotReq{}
	req.SenderEmail = m.AdminEmail()
	req.Type = "text"
	req.Message = message
	req.RoomID = roomID
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), nil)
	r.AddHeader("Authorization", m.AdminToken())
	err := r.DoRequest()

	return err
}

// SetToogleBotInRoom set tootle bot in room
func (m *MultichannelImpl) SetToogleBotInRoom(roomID string, isActive bool) (*SetToogleBotInRoomResponse, *qiscus.Error) {
	resp := &SetToogleBotInRoomResponse{}
	url := fmt.Sprintf("%s/bot/%s/activate", m.APIBase(), roomID)

	req := SetToogleBotInRoomReq{IsActive: isActive}
	jsonReq, _ := json.Marshal(req)

	r := qiscus.NewHttpRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq), resp)
	r.AddHeader("Authorization", m.AdminToken())
	err := r.DoRequest()

	return resp, err
}
