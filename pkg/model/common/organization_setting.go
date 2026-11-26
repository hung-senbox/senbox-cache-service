package common

import "time"

type OrganizationSetting struct {
	DateCount  int       `json:"date_count"`
	MediaCount int       `json:"media_count"`
	CountDown  CountDown `json:"count_down"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
}

type CountDown struct {
	DateCount  int `json:"date_count"`
	MediaCount int `json:"media_count"`
}
