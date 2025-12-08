package models

type ChatMessage struct {
	SessionID    string `json:"session_id"`
	MessageIndex int    `json:"message_index"`
	Sender       string `json:"sender"`
	Content      string `json:"content"`
	TokensUsed   int    `json:"tokens_used"`
	ModelVersion string `json:"model_version"`
	CreatedAt    string `json:"created_at"` //ISO string
}
