package models

type User struct {
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	Picture     string `json:"picture"`
	CreatedAt   string `json:"created_at"`
}
