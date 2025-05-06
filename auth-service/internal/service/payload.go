package service

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsValid  bool   `json:"is_valid"`
}
