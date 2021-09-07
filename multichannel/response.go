package multichannel

import "time"

// RoomTagsResponse : Represent Get room tags response payload
type RoomTagsResponse struct {
	Data []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

// CreateRoomTagResponse: Represent Create room tag response payload
type CreateRoomTagResponse struct {
	Data struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

// CreateAdditionalInfoRoomResponse: Represents Create additional info room response payload
type CreateAdditionalInfoRoomResponse struct {
	Data struct {
		Extras struct {
			UserProperties []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"user_properties"`
		} `json:"extras"`
		FirstInitiated struct {
			Date         string `json:"date"`
			TimezoneType int    `json:"timezone_type"`
			Timezone     string `json:"timezone"`
		} `json:"first_initiated"`
		FirstAgentResponseTime struct {
			Date         string `json:"date"`
			TimezoneType int    `json:"timezone_type"`
			Timezone     string `json:"timezone"`
		} `json:"first_agent_response_time"`
		UserID string `json:"user_id"`
	} `json:"data"`
}

// GetAdditionalInfoRoomResponse: Represent Get additional info room response payload
type GetAdditionalInfoRoomResponse struct {
	Data struct {
		Extras struct {
			UserProperties []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"user_properties"`
		} `json:"extras"`
		FirstInitiated         time.Time `json:"first_initiated"`
		FirstAgentResponseTime time.Time `json:"first_agent_response_time"`
		UserID                 string    `json:"user_id"`
		ChannelID              int       `json:"channel_id"`
		IsBlocked              bool      `json:"is_blocked"`
		ChannelName            string    `json:"channel_name"`
		Channel                struct {
			ID                  int         `json:"id"`
			AppCode             string      `json:"app_code"`
			SecretKey           string      `json:"secret_key"`
			CreatedAt           string      `json:"created_at"`
			UpdatedAt           string      `json:"updated_at"`
			IsActive            bool        `json:"is_active"`
			AppID               int         `json:"app_id"`
			ForwardURL          interface{} `json:"forward_url"`
			ForwardEnabled      bool        `json:"forward_enabled"`
			Name                string      `json:"name"`
			BadgeURL            string      `json:"badge_url"`
			UseChannelResponder bool        `json:"use_channel_responder"`
		} `json:"channel"`
	} `json:"data"`
}

// SetToogleBotInRoomResponse: Represent Set toogle bot in room response payload
type SetToogleBotInRoomResponse struct {
	Data struct {
		ID                    int         `json:"id"`
		AppID                 int         `json:"app_id"`
		UserID                string      `json:"user_id"`
		RoomID                string      `json:"room_id"`
		Source                string      `json:"source"`
		CreatedAt             string      `json:"created_at"`
		UpdatedAt             string      `json:"updated_at"`
		IsHandledByBot        bool        `json:"is_handled_by_bot"`
		StartServiceCommentID string      `json:"start_service_comment_id"`
		UserAvatarURL         string      `json:"user_avatar_url"`
		Name                  string      `json:"name"`
		HasNoMessage          bool        `json:"has_no_message"`
		Extras                string      `json:"extras"`
		CheckWaContact        bool        `json:"check_wa_contact"`
		Origin                string      `json:"origin"`
		RoomBadge             string      `json:"room_badge"`
		IsWaiting             bool        `json:"is_waiting"`
		SubSource             interface{} `json:"sub_source"`
		ChannelID             int         `json:"channel_id"`
		Resolved              bool        `json:"resolved"`
		ResolvedTs            interface{} `json:"resolved_ts"`
		Type                  interface{} `json:"type"`
		DeletedAt             interface{} `json:"deleted_at"`
		CustomerID            interface{} `json:"customer_id"`
	} `json:"data"`
}

// LoginAdminResponse: Represent Login admin response payload
type LoginAdminResponse struct {
	Data struct {
		User struct {
			ID                  int           `json:"id"`
			Name                string        `json:"name"`
			Email               string        `json:"email"`
			AuthenticationToken string        `json:"authentication_token"`
			CreatedAt           string        `json:"created_at"`
			UpdatedAt           string        `json:"updated_at"`
			SdkEmail            string        `json:"sdk_email"`
			SdkKey              string        `json:"sdk_key"`
			IsAvailable         bool          `json:"is_available"`
			Type                int           `json:"type"`
			AvatarURL           string        `json:"avatar_url"`
			AppID               int           `json:"app_id"`
			IsVerified          bool          `json:"is_verified"`
			NotificationsRoomID string        `json:"notifications_room_id"`
			BubbleColor         interface{}   `json:"bubble_color"`
			QismoKey            string        `json:"qismo_key"`
			DirectLoginToken    interface{}   `json:"direct_login_token"`
			LastLogin           string        `json:"last_login"`
			ForceOffline        bool          `json:"force_offline"`
			DeletedAt           interface{}   `json:"deleted_at"`
			TypeAsString        string        `json:"type_as_string"`
			AssignedRules       []interface{} `json:"assigned_rules"`
			App                 struct {
				ID                             int         `json:"id"`
				Name                           string      `json:"name"`
				AppCode                        string      `json:"app_code"`
				SecretKey                      string      `json:"secret_key"`
				CreatedAt                      string      `json:"created_at"`
				UpdatedAt                      string      `json:"updated_at"`
				BotWebhookURL                  string      `json:"bot_webhook_url"`
				IsBotEnabled                   bool        `json:"is_bot_enabled"`
				AllocateAgentWebhookURL        string      `json:"allocate_agent_webhook_url"`
				IsAllocateAgentWebhookEnabled  bool        `json:"is_allocate_agent_webhook_enabled"`
				MarkAsResolvedWebhookURL       string      `json:"mark_as_resolved_webhook_url"`
				IsMarkAsResolvedWebhookEnabled bool        `json:"is_mark_as_resolved_webhook_enabled"`
				IsMobilePnEnabled              bool        `json:"is_mobile_pn_enabled"`
				IsActive                       bool        `json:"is_active"`
				IsSessional                    bool        `json:"is_sessional"`
				IsAgentAllocationEnabled       bool        `json:"is_agent_allocation_enabled"`
				IsAgentTakeoverEnabled         bool        `json:"is_agent_takeover_enabled"`
				IsTokenExpiring                bool        `json:"is_token_expiring"`
				PaidChannelApproved            interface{} `json:"paid_channel_approved"`
				UseLatest                      bool        `json:"use_latest"`
				AppConfig                      struct {
					ID                     int         `json:"id"`
					AppID                  int         `json:"app_id"`
					Widget                 string      `json:"widget"`
					CreatedAt              string      `json:"created_at"`
					UpdatedAt              string      `json:"updated_at"`
					OfflineMessage         interface{} `json:"offline_message"`
					OnlineMessage          string      `json:"online_message"`
					Timezone               string      `json:"timezone"`
					EnableBulkAssign       bool        `json:"enable_bulk_assign"`
					SendOnlineIfResolved   bool        `json:"send_online_if_resolved"`
					SendOfflineEachMessage bool        `json:"send_offline_each_message"`
				} `json:"app_config"`
				AgentRoles []struct {
					ID            int    `json:"id"`
					AppID         int    `json:"app_id"`
					Name          string `json:"name"`
					IsDefaultRole bool   `json:"is_default_role"`
					CreatedAt     string `json:"created_at"`
					UpdatedAt     string `json:"updated_at"`
				} `json:"agent_roles"`
			} `json:"app"`
		} `json:"user"`
		Details struct {
			IsIntegrated bool `json:"is_integrated"`
			SdkUser      struct {
				ID          int    `json:"id"`
				Token       string `json:"token"`
				Email       string `json:"email"`
				Password    string `json:"password"`
				DisplayName string `json:"display_name"`
				AvatarURL   string `json:"avatar_url"`
				Extras      struct {
					Type            string      `json:"type"`
					UserBubbleColor interface{} `json:"user_bubble_color"`
				} `json:"extras"`
			} `json:"sdk_user"`
			App struct {
				AppCode                        string `json:"app_code"`
				SecretKey                      string `json:"secret_key"`
				Name                           string `json:"name"`
				BotWebhookURL                  string `json:"bot_webhook_url"`
				IsBotEnabled                   bool   `json:"is_bot_enabled"`
				IsAllocateAgentWebhookEnabled  bool   `json:"is_allocate_agent_webhook_enabled"`
				AllocateAgentWebhookURL        string `json:"allocate_agent_webhook_url"`
				MarkAsResolvedWebhookURL       string `json:"mark_as_resolved_webhook_url"`
				IsMarkAsResolvedWebhookEnabled bool   `json:"is_mark_as_resolved_webhook_enabled"`
				IsActive                       bool   `json:"is_active"`
				IsSessional                    bool   `json:"is_sessional"`
				IsAgentAllocationEnabled       bool   `json:"is_agent_allocation_enabled"`
				IsAgentTakeoverEnabled         bool   `json:"is_agent_takeover_enabled"`
				UseLatest                      bool   `json:"use_latest"`
				IsBulkAssignmentEnabled        bool   `json:"is_bulk_assignment_enabled"`
			} `json:"app"`
		} `json:"details"`
		LongLivedToken string `json:"long_lived_token"`
		UserConfigs    struct {
			Notifagentjoining           interface{} `json:"notifagentjoining"`
			IsNotifagentjoiningEnabled  bool        `json:"is_notifagentjoining_enabled"`
			Notifmessagecoming          interface{} `json:"notifmessagecoming"`
			IsNotifmessagecomingEnabled bool        `json:"is_notifmessagecoming_enabled"`
		} `json:"user_configs"`
	} `json:"data"`
}
