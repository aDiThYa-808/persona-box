package models

type ChatSession struct {
	SessionID    string `json:"session_id"`
	PersonaID    string `json:"persona_id"`
	Title        string `json:"title"`
	CreatedAt    string `json:"created_at"` //ISO string
	UpdatedAt    string `json:"updated_at"` //ISO string
	MessageCount int    `json:"message_count"`
	TokenCount   int    `json:"token_count"`
	Summary      string `json:"summary"`
}
