package router

import (
	"net/http"

	"github.com/chitano/chatapp/internal/auth/handler"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(authHandler *handler.AuthHandler) {
	r = gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to ChatAPP"})
	})

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
}

func StartApp(addr string) error {
	return r.Run(addr)
}
