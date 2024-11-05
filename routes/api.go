package routes

import (
	"github.com/goravel/framework/facades"
	"goravel/app/http/middleware"

	"goravel/app/http/controllers"
)

func Api() {
	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	facades.Route().Post("/login", authController.Login)
	facades.Route().Post("/logout", authController.Logout)
	facades.Route().Put("/register", authController.Register)
	facades.Route().Middleware(middleware.AuthMiddleware()).Post("users", userController.Update)
}
