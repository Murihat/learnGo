package services

import (
	"learnGo/models"
	"learnGo/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.CategoryModel, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *models.CategoryModel) error {
	return s.repo.Create(data)
}

func (s *CategoryService) GetByID(id int) (*models.CategoryModel, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(category *models.CategoryModel) error {
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
