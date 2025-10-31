package mainservicedto

import "time"

type UserResponse struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Fullname     string    `json:"fullname"`
	Nickname     string    `json:"nickname"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Birthday     time.Time `json:"birthday"`
	QRLogin      string    `json:"qr_login"`
	Avatar       string    `json:"avatar"`
	AvatarURL    string    `json:"avatar_url"`
	Password     string    `json:"password"`
	IsBlocked    bool      `json:"is_blocked"`
	BlockedAt    time.Time `json:"blocked_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CustomID     string    `json:"custom_id"`
	ReLoginWeb   bool      `json:"re_login_web"`
	CreatedIndex int       `json:"created_index"`
}
