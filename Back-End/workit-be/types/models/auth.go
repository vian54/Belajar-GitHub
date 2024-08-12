package models

type AuthResponse struct {
	Token     string `json:"token"`
	User      User   `json:"user"`
	Role      Role   `json:"role"`
	ExpiredAt string `json:"expired_at"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
