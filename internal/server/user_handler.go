package server

import (
	"github.com/davidbassouto/ecommerce-golang/internal/dto"
	"github.com/davidbassouto/ecommerce-golang/internal/services"
	"github.com/davidbassouto/ecommerce-golang/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) getProfile(c *gin.Context) {

	userID := c.GetUint("user_id")
	userService := services.NewUserService(s.db)
	responseProfile, err := userService.GetProfile(userID)
	if err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}
	utils.SuccessResponse(c, "Profile retrieved successfully", responseProfile)
}

func (s *Server) updateProfile(c *gin.Context) {
	userID := c.GetUint("user_id") // middlewares

	var req dto.UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return

	}
	userService := services.NewUserService(s.db)
	responseProfile, err := userService.UpdateProfile(userID, &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update profile", err)
		return
	}

	utils.SuccessResponse(c, "Profile updated successfully", responseProfile)

}
