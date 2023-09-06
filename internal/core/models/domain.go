package models

type Domain struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BlockedBy   string `json:"blockedBy"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
