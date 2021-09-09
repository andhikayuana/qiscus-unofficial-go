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

// GetAllAgentsReq: Represent Get all agents request payload
type GetAllAgentsReq struct {
	// Page: default 1
	Page int

	// Limit: default 20
	Limit  int
	Search string

	// Scope: either `division`, `name`, or `email`, or default
	Scope string
}

// AssignAgentReq: Represent Assign agent request payload
type AssignAgentReq struct {
	RoomID             string `json:"room_id"`
	AgentID            string `json:"agent_id"`
	ReplaceLatestAgent bool   `json:"replace_latest_agent"`

	// MaxAgent: default max agent is 5
	MaxAgent int `json:"max_agent"`
}

// GetAgentsByDivisionReq: Represent Get agents by division request payload
type GetAgentsByDivisionReq struct {
	// Page: default 1
	Page int

	// Limit: default 20
	Limit       int
	DivisionIDs []string

	// IsAvailable: online availability filter, default all, can be true of false
	IsAvailable bool

	// Sort: default asc (less customer count) can be desc
	Sort string
}

// GetAllDivisionReq: Represent Get all division request payload
type GetAllDivisionReq struct {
	// Page: default 1
	Page int

	// Limit: default 20
	Limit int
}

// MarkAsResolvedReq: Represent Mark as resolved request payload
type MarkAsResolvedReq struct {
	RoomID        string `json:"room_id"`
	Notes         string `json:"notes"`
	LastCommentID string `json:"last_comment_id"`
}
