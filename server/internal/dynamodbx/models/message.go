package models

type ChatMessage struct {
	SessionID string `json:"session_id"` //PK
	CreatedAt string `json:"created_at"` //SK ISO string
	Sender    string `json:"sender"`
	Content   string `json:"content"`
}
