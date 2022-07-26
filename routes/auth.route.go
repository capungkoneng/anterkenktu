package routes

import (
	"github.com/capungkoneng/anterkenktu/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")
	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.LoginUser)
}
