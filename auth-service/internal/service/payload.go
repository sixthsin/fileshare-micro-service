package service

type UserResponse struct {
	ID      string `json:"user_id"`
	Email   string `json:"email"`
	IsValid bool   `json:"is_valid"`
}
