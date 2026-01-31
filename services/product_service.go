package services

import (
	"errors"
	"learnGo/models"
	"learnGo/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]models.ProductModel, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(data *models.ProductModel) error {
	if data.CategoryID == 0 {
		return errors.New("category_id is required")
	}
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id int) (*models.ProductModel, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.ProductModel) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
