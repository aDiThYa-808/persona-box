package models

// Model for User table
type User struct {
	//UserID is the PARTITION KEY of User table
	UserID string `json:"user_id"`

	Email       string `json:"email"`
	DisplayName string `json:"display_name"`

	//CreatedAt must be an ISO string. Format: RFC3339
	CreatedAt string `json:"created_at"`
}
