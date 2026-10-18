package handlers

import (
	"backend/internal/dto"
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}
func (h *AuthHandler) Register(c *gin.Context) {
	var input dto.RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.AuthService.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user":    user,
	})
}
