package sdk

// LoginOrRegisterReq: Represent Login or register request payload
type LoginOrRegisterReq struct {
	UserID    string `json:"user_id"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// ResetUserTokenReq: Represent Reset user token request payload
type ResetUserTokenReq struct {
	UserID string `json:"user_id"`
}

// CreateRoomReq: Represent Create room request payload
type CreateRoomReq struct {
	RoomName      string   `json:"room_name"`
	Creator       string   `json:"creator"`
	Participants  []string `json:"participants"`
	RoomAvatarURL string   `json:"room_avatar_url"`
	RoomOptions   string   `json:"room_options"`
}

// GetOrCreateRoomWithTargetReq: Represent Get or create room with target request payload
type GetOrCreateRoomWithTargetReq struct {
	UserIDs     []string `json:"user_ids"`
	RoomOptions string   `json:"room_options"`
}

// UpdateRoomReq: Represent Update room request payload
type UpdateRoomReq struct {
	RoomID      string `json:"room_id"`
	RoomName    string `json:"room_name"`
	RoomOptions string `json:"room_options"`
}

// AddRoomParticipants: Represent Add room participants request payload
type AddRoomParticipantsReq struct {
	RoomID  string   `json:"room_id"`
	UserIds []string `json:"user_ids"`
}

// RemoveRoomParticipantsReq: Represent Remove room participants request payload
type RemoveRoomParticipantsReq struct {
	RoomID  string   `json:"room_id"`
	UserIds []string `json:"user_ids"`
}

// PostCommentReq: Represent Post comment request payload
type PostCommentReq struct {
	UserID  string      `json:"user_id"`
	RoomID  string      `json:"room_id"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Extras  interface{} `json:"extras"`
	Payload interface{} `json:"payload"`
}
