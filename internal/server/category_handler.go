package server

import (
	"strconv"

	"github.com/davidbassouto/ecommerce-golang/internal/dto"
	"github.com/davidbassouto/ecommerce-golang/internal/services"
	"github.com/davidbassouto/ecommerce-golang/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) createCategory(c *gin.Context) {
	// atribuir req
	// vincular JSON
	// criar chamada para service

	var req dto.CreateCategoryRequest // ja cria instancia e seta valores como nil
	// (req := dto.CreateCategoryRequest sozinho é apenas um tipo, não cria uma instância)

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid Request data", err)
		return
	}

	categoryService := services.NewCategoryService(s.db)
	category, err := categoryService.CreateCategory(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create category", err)
	}
	utils.CreatedResponse(c, "Category created successfully", category)
}

func (s *Server) getCategories(c *gin.Context) {
	categoryService := services.NewCategoryService(s.db)
	categories, err := categoryService.GetCategories()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch categories", err)
	}

	utils.SuccessResponse(c, "Categories retrieved successfully", categories)
}

func (s *Server) updateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid category ID", err)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid Request data", err)
		return
	}
	categoryService := services.NewCategoryService(s.db)
	category, err := categoryService.UpdateCategory(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update category", err)
		return
	}
	utils.SuccessResponse(c, "Category updated successfully", category)
}

func (s *Server) deleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid category ID", err)
		return
	}

	categoryService := services.NewCategoryService(s.db)
	if err := categoryService.DeleteCategory(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete category", err)
		return
	}

	utils.SuccessResponse(c, "Category deleted successfully", nil)
}
