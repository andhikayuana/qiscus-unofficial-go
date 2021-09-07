package multichannel

// CreateRoomTagReq: Represent Create room tag request payload
type CreateRoomTagReq struct {
	RoomID string `json:"room_id"`
	Tag    string `json:"tag"`
}

type UserProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CreateAdditionalInfoRoomReq: Represent Create additional room request payload
type CreateAdditionalInfoRoomReq struct {
	UserProperties []UserProperty `json:"user_properties"`
}

// SendMessageTextByBOTReq: Represent Send message text by Bot request payload
type SendMessageTextByBotReq struct {
	SenderEmail string `json:"sender_email"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	RoomID      string `json:"room_id"`
}

// SetToogleBotInRoomReq: Represent Set toogle room request payload
type SetToogleBotInRoomReq struct {
	IsActive bool `json:"is_active"`
}

// LoginAdminReq: Represent Login admin request payload
type LoginAdminReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
