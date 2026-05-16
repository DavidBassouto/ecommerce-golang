package services

// CRUD COMPLETO COM NOTES
import (
	"github.com/davidbassouto/ecommerce-golang/internal/dto"
	"github.com/davidbassouto/ecommerce-golang/internal/models"
	"github.com/davidbassouto/ecommerce-golang/internal/utils"
	"gorm.io/gorm"
)

// encapsula uma conexão com o banco de dados
// (usando a biblioteca GORM) para realizar operações
type ProductService struct {
	db *gorm.DB
}

// NewProductService é um construtor que cria uma nova instância de ProductService
// Recebe uma conexão de banco de dados (db) e retorna um ponteiro para ProductService
// O campo db do ProductService é inicializado com a conexão recebida
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) convertToProductResponse(product *models.Product) dto.ProductResponse {
	images := make([]dto.ProductImageResponse, len(product.Images))
	for i := range product.Images {
		images[i] = dto.ProductImageResponse{
			ID:        product.Images[i].ID,
			URL:       product.Images[i].URL,
			AltText:   product.Images[i].AltText,
			IsPrimary: product.Images[i].IsPrimary,
		}
	}
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category: dto.CategoryResponse{
			ID:          product.CategoryID,
			Name:        product.Category.Name,
			Description: product.Category.Description,
			IsActive:    product.Category.IsActive,
		},
		IsActive:  product.IsActive,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Images:    images,
	}
}

func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		SKU:         req.SKU,
	}

	if err := s.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ProductService) GetProducts(page, limit int) ([]dto.ProductResponse, *utils.PaginationMeta, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	// cria um slice nulo (nil), e o próprio GORM se encarregará de alocar a memória necessária
	// quando encontrar os registros no banco.
	var products []models.Product
	var total int64

	// verificar qnts produtos estao ativos
	// Model indica a tabela alvo (semelhante ao FROM em SQL)
	s.db.Model(&models.Product{}).Where("is_active=?", true).Count(&total)

	// usar preload para carregar sub arrays do objeto (nesse caso, categorias e images)
	// preload e tipo um join
	// carregar as outras classes como parte do item
	if err := s.db.Preload("Category").Preload("Images").Where("is_active=?", true).Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, nil, err
	}

	// transformar em uma resposta
	// loop usando o make pra ja criar a quantidade de posições
	// make (tipo, tamanho do slice)

	response := make([]dto.ProductResponse, len(products))

	// loop com for i := range slice a ser iterado
	// cada item em products vai ser alocado a uma posicao ja criada em response
	for i := range products {
		response[i] = s.convertToProductResponse(&products[i])
	}

	// calculate a quantidade de paginas
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	meta := &utils.PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
	return response, meta, nil
}

func (s *ProductService) GetProductByID(id uint) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.Preload("Category").Preload("Images").Where("id=?", id).First(&product).Error; err != nil {
		return nil, err
	}

	response := s.convertToProductResponse(&product)
	return &response, nil

}

func (s *ProductService) UpdateProductByID(id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	var product models.Product

	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	product.CategoryID = req.CategoryID
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	if err := s.db.Save(&product).Error; err != nil {
		return nil, err
	}

	return s.GetProductByID(id)
}

func (s *ProductService) DeleteProductByID(id uint) error {
	return s.db.Delete(&models.Product{}, id).Error
}
