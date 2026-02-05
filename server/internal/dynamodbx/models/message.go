package models

type ChatMessage struct {
	SessionID string `json:"session_id"` //PK
	CreatedAt string `json:"created_at"` //SK ISO string
	Role      string `json:"role"`
	Message   string `json:"message"`
}
