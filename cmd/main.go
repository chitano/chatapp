package main

import (
	"os"

	"github.com/chitano/chatapp/config"
	"github.com/chitano/chatapp/internal/auth/handler"
	"github.com/chitano/chatapp/internal/auth/services"
	"github.com/chitano/chatapp/internal/user/repositories"
	"github.com/chitano/chatapp/router"
)

func main() {
	config.ConnectDatabase(config.LoadConfig())
	userRepo := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	router.InitRouter(authHandler)
	router.StartApp(os.Getenv("APP_URL"))
}
