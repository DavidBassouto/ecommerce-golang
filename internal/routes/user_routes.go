package routes

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(
	api *gin.RouterGroup,
	authMiddleware gin.HandlerFunc,
	profile gin.HandlerFunc,
	updateProfile gin.HandlerFunc,
) {
	user := api.Group("users")
	user.Use(authMiddleware)
	{
		user.GET("profile", profile)
		user.PUT("profile", updateProfile)
	}
}
