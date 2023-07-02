package models

// Domain is used to represent a blocked domain
type Domain struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BlockedBy   string `json:"blockedBy"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// the user's details
type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
}
