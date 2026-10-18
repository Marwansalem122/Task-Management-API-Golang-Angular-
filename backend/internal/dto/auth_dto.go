package dto

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}
type AuthRequest struct {
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	User    interface{} `json:"user,omitempty"`
}
