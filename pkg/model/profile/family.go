package profile

import "time"

type Family struct {
	ID        string         `json:"id"`
	Father    FamilyMember   `json:"father"`
	Mother    FamilyMember   `json:"mother"`
	Children  []FamilyMember `json:"children"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type FamilyMember struct {
	ID string `json:"id"`
}
