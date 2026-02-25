package models

// Model for ChatMessage table
type ChatMessage struct {
	// SessionID is the PARTITION KEY of ChatMessage table
	SessionID string `json:"session_id"`

	// CreatedAt is the SORT KEY of ChatMessage table.
	// CreatedAt must be an ISO string. Format: RFC3339
	CreatedAt string `json:"created_at"` //SK ISO string

	// Role can be either user or assistant
	Role string `json:"role"`

	Message string `json:"message"`
}
