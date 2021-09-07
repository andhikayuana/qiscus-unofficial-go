package sdk

import "time"

// LoginOrRegisterResponse: Represent Login or register response payload
type LoginOrRegisterResponse struct {
	Results struct {
		User struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetUserProfileResponse: Represent Get user profile response payload
type GetUserProfileResponse struct {
	Results struct {
		User struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetUserTokenResponse: Represent Get user token response payload
type GetUserTokenResponse struct {
	Results struct {
		Token string `json:"token"`
	} `json:"results"`
	Status int `json:"status"`
}

// CreateRoomResponse: Represent Create room response payload
type CreateRoomResponse struct {
	Results struct {
		Room struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"room"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetRoomsInfoResponse: Represent Get rooms info response payload
type GetRoomsInfoResponse struct {
	Results struct {
		Rooms []struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"rooms"`
	} `json:"results"`
	Status int `json:"status"`
}

// UpdateRoomResponse: Represent Update room response payload
type UpdateRoomResponse struct {
	Results struct {
		Changed bool `json:"changed"`
		Room    struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"room"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetRoomParticipantsResponse: Represent Get room participants response payload
type GetRoomParticipantsResponse struct {
	Results struct {
		Participants []struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"participants"`
	} `json:"results"`
	Status int `json:"status"`
}

// AddRoomParticipantsResponse: Represent Add room participants response payload
type AddRoomParticipantsResponse struct {
	Results struct {
		ParticipantsAdded []struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"participants_added"`
	} `json:"results"`
	Status int `json:"status"`
}

// RemoveRoomParticipantsResponse: Represent Remove room participants response payload
type RemoveRoomParticipantsResponse struct {
	Results struct {
		ParticipantsRemoved []struct {
			Active    bool   `json:"active"`
			AvatarURL string `json:"avatar_url"`
			Extras    struct {
			} `json:"extras"`
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		} `json:"participants_removed"`
	} `json:"results"`
	Status int `json:"status"`
}

// GetUserRoomsResponse: Represent Get user rooms response payload
type GetUserRoomsResponse struct {
	Results struct {
		Meta struct {
			CurrentPage int `json:"current_page"`
			TotalRoom   int `json:"total_room"`
		} `json:"meta"`
		Rooms []struct {
			RoomAvatarURL string `json:"room_avatar_url"`
			RoomChannelID string `json:"room_channel_id"`
			RoomID        string `json:"room_id"`
			RoomName      string `json:"room_name"`
			RoomOptions   string `json:"room_options"`
			RoomType      string `json:"room_type"`
		} `json:"rooms"`
	} `json:"results"`
	Status int `json:"status"`
}

// PostCommentResponse: Represent Post comment response payload
type PostCommentResponse struct {
	Results struct {
		Comment struct {
			Extras struct {
			} `json:"extras"`
			ID      int    `json:"id"`
			Message string `json:"message"`
			Payload struct {
			} `json:"payload"`
			Timestamp time.Time `json:"timestamp"`
			Type      string    `json:"type"`
			User      struct {
				Active    bool   `json:"active"`
				AvatarURL string `json:"avatar_url"`
				Extras    struct {
				} `json:"extras"`
				UserID   string `json:"user_id"`
				Username string `json:"username"`
			} `json:"user"`
		} `json:"comment"`
	} `json:"results"`
	Status int `json:"status"`
}
