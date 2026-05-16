package server

import (
	"strconv"

	"github.com/davidbassouto/ecommerce-golang/internal/dto"
	"github.com/davidbassouto/ecommerce-golang/internal/services"
	"github.com/davidbassouto/ecommerce-golang/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) createProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.CreateProduct(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Creating product failed", err)
		return
	}
	utils.CreatedResponse(c, "Product created successfully", product)

}

func (s *Server) getProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	productService := services.NewProductService(s.db)
	products, meta, err := productService.GetProducts(page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Fetching products failed", err)
		return
	}

	utils.PaginatedSuccessResponse(c, "Products retrieved successfully", products, *meta)

}

func (s *Server) getProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID", err)
		return
	}

	productService := services.NewProductService(s.db)
	product, err := productService.GetProductByID(uint(id))
	if err != nil {
		utils.InternalServerErrorResponse(c, "Product not found", err)
	}
	utils.SuccessResponse(c, "Product retrieved successfully", product)
}

func (s *Server) updateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID", err)
		return
	}

	// update tem que fazer verificação de shouldbindjson

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
	}

	productService := services.NewProductService(s.db)
	product, err := productService.UpdateProductByID(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update product", err)
	}
	utils.SuccessResponse(c, "Product updated successfully", product)

}

func (s *Server) deleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID", err)
		return
	}

	productService := services.NewProductService(s.db)
	if err := productService.DeleteProductByID(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete product", err)
		return
	}

	utils.SuccessResponse(c, "Product deleted successfully", nil)
}
