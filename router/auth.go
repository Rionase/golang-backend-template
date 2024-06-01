package router

import (
	"golang-backend-template/controller"
	"golang-backend-template/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController controller.IAuthController) {
	authRoute := router.Group("/auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register/user", authController.RegisterUser)
		authRoute.POST("/register/admin", middleware.WithAuth("admin"), authController.RegisterAdmin)
		authRoute.POST("/logout", middleware.WithAuth("user", "admin"), authController.Logout)
	}
}
