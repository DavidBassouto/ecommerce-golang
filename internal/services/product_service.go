package services

// CRUD COMPLETO COM NOTES
import (
	"github.com/davidbassouto/ecommerce-golang/internal/dto"
	"github.com/davidbassouto/ecommerce-golang/internal/models"
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
