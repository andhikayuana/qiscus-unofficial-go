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

// GetRoomParticipantsReq: Represent Get room participant request payload
type GetRoomParticipantsReq struct {
	RoomID string
	Page   int // default 1
	Limit  int // default 20
}

// GetUserRoomsReq: Represent Get user room request payload
type GetUserRoomsReq struct {
	UserID string
	Page   int // default 1
	Limit  int // default 20
}

// LoadCommentsReq: Represent Load comments request payload
type LoadCommentsReq struct {
	RoomID string
	Page   int // default 1
	Limit  int // default 20
}

// PostSystemEventMessageReq: Represent Post system event message request payload
type PostSystemEventMessageReq struct {
	RoomID  string
	Message string
	Payload interface{}
	Extras  interface{}
}

// GetUnreadCountReq: Represent Get unread count request payload
type GetUnreadCountReq struct {
	UserID  string
	RoomIDs []string
}

// GetUsersReq: Represent Get users request payload
type GetUsersReq struct {
	Page       int    // default 1
	Limit      int    // default 20
	ShowAll    bool   // default false
	OrderQuery string // default 'created_at desc nulls last'
}

// LoadCommentsWithRangeReq: Represent Load comments with range request payload
type LoadCommentsWithRangeReq struct {
	RoomID         string
	FirstCommentID string
	LastCommentID  string
}

// GetOrCreateChannelReq: Represent Get or create channel request payload
type GetOrCreateChannelReq struct {
	UniqueID      string   `json:"unique_id"`
	RoomName      string   `json:"room_name"`
	Participants  []string `json:"participants"`
	RoomAvatarURL string   `json:"room_avatar_url"`
	RoomOptions   string   `json:"room_options"`
}

// GetAverageReplyTimeUserReq: Represent Get average reply time user request payload
type GetUserResponseRateReq struct {
	UserIDs   []string
	StartTime string // in format "YYYY-MM-DD hh:mm:ss"
	EndTime   string // in format "YYYY-MM-DD hh:mm:ss"
}

// GetAverageReplyTimeUserReq: Represent Get average reply time user request payload
type GetAverageReplyTimeUserReq struct {
	UserID    string
	StartTime string // in format "YYYY-MM-DD hh:mm:ss"
	EndTime   string // in format "YYYY-MM-DD hh:mm:ss"
}

// GetWebhookLogsReq: Represent Get webhook logs request payload
type GetWebhookLogsReq struct {
	Page  int    // default 1, max 100
	Limit int    // default 20,max 100
	Type  string // default all, can be 'mobile' or 'rest'
}

// DeactivateUserReq: Represent Deactivate user request payload
type DeactivateUserReq struct {
	UserIDs []string `json:"user_ids"`
}

// ReactivateUserReq: Represent Reactivate user request payload
type ReactivateUserReq struct {
	UserIDs []string `json:"user_ids"`
}
