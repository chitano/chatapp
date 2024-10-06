package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type AppJWTClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8"`
}

type AuthUser struct {
	ID       uint64 `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type LoginResponse struct {
	AccessToken string    `json:"access_token,omitempty"`
	ExpiresAt   int64     `json:"expires_at"`
	AuthUser    *AuthUser `json:"user,omitempty"`
}

type RegisterRequest struct {
	Name     string `form:"name" binding:"required"`
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required,min=8"`
	Contact  string `form:"contact" binding:"required"`
}
