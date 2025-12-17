package models

type User struct {
	UserID      string `json:"user_id"` // PK
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	CreatedAt   string `json:"created_at"`
}
