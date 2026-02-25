package models

// Model for ChatSession
type ChatSession struct {
	//SessionID is the SORT KEY of ChatSession table
	SessionID string `json:"session_id"`

	//PersonaID is the PARTITION KEY of ChatSession table
	PersonaID string `json:"persona_id"`

	Title string `json:"title"`

	//CreatedAt must be an ISO string. Format: RFC3339
	CreatedAt string `json:"created_at"`

	//UpdatedAt must be an ISO string. Format: RFC3339
	UpdatedAt string `json:"updated_at"`

	MessageCount int    `json:"message_count"`
	TokenCount   int    `json:"token_count"`
	Summary      string `json:"summary"`
}

// Type ChatSessionList is a subset of Type ChatSession, safe for API response
type ChatSessionList struct {
	SessionID string `json:"session_id"`
	Title     string `json:"title"`

	//UpdatedAt must be an ISO string. Format: RFC3339
	UpdatedAt string `json:"updated_at"`
}
