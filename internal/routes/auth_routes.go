package routes

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(
	api *gin.RouterGroup,
	register gin.HandlerFunc,
	login gin.HandlerFunc,
	logout gin.HandlerFunc,
	refreshToken gin.HandlerFunc,
) {
	auth := api.Group("auth")
	{
		auth.POST("register", register)
		auth.POST("login", login)
		auth.POST("logout", logout)
		auth.POST("refresh", refreshToken)
	}
}
