package services

// CRUD COMPLETO COM NOTES
import (
	"github.com/davidbassouto/ecommerce-golang/internal/dto"
	"github.com/davidbassouto/ecommerce-golang/internal/models"
	"github.com/davidbassouto/ecommerce-golang/internal/utils"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	// modelar para o banco, por isso importar model
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	// tentar gravar no database
	// chamar create
	if err := s.db.Create(&category).Error; err != nil {
		return nil, err
	}
	// se nao deu erro, retorna o DTO

	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

func (s *ProductService) GetCategories() ([]dto.CategoryResponse, error) {
	// criar um slice de categorias
	var categories []models.Category
	// buscar no banco com .Where pra adicionar condicao e .Find pra percorer todos os registros e preencher o slice
	if err := s.db.Where("is_active = ?", true).Find(&categories).Error; err != nil {
		return nil, err
	}
	// se nao deu erro montar o response com o make + DTO + len
	// make aloca espaço na memoria, len diz quants espaços

	response := make([]dto.CategoryResponse, len(categories))

	// necessario um loop para preencher "gavetas" que o make criou
	for i := range categories {
		response[i] = dto.CategoryResponse{
			ID:          categories[i].ID,
			Name:        categories[i].Name,
			Description: categories[i].Description,
			IsActive:    categories[i].IsActive,
		}
	}
	return response, nil
}

func (s *ProductService) UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	// reatribuir os valores:
	category.Name = req.Name
	category.Description = req.Description
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	// tentar salvar e ja retornar um erro se nao funcionar
	// tx := s.db.Save(&category) // tx é do tipo *gorm.DB
	// err := tx.Error            // pegamos o erro de dentro do objeto
	// if err != nil {            // verificamos se ele existe
	// 	return nil, err
	// }
	if err := s.db.Save(&category).Error; err != nil {
		return nil, err
	}
	return &dto.CategoryResponse{
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

func (s *ProductService) DeleteCategory(id uint) error {
	// so uma linha que ja cria a variavel Category e deleta ela via id
	return s.db.Delete(&models.Category{}, id).Error
}

// PRODUCTS

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

	var products []models.Product
	var total int64

	// verificar qnts produtos estao ativos
	s.db.Model(&models.Product{}).Where("is_active=?", true).Count(&total)

	// usar preload para carregar sub arrays do objeto (nesse caso, categorias e images)
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
