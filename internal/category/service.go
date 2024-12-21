package category

import (
	"strings"
	"trackergo/internal/users"

	"github.com/google/uuid"
)

type CategoryService interface {
	CreateDefaultCategories() error
	CreateCategory(userId uuid.UUID, name string) (*Category, error)
	GetCategoryById(categoryId, userId uuid.UUID) (*Category, error)
	UpdateCategory(categoryId, userId uuid.UUID, name string) (*Category, error)
	DeleteCategory(categoryId, userId uuid.UUID) error
	GetAllCategories(userId uuid.UUID) ([]Category, error)
}

type categoryService struct {
	repo CategoryRepository
}

var (
  _ CategoryService = &categoryService{}
)

func NewCategoryService(repo CategoryRepository) *categoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateDefaultCategories() error {
	err := s.repo.CreateDefaultCategories()
	if err != nil {
		return err
	}
	return nil
}

func (s *categoryService) CreateCategory(userId uuid.UUID, name string) (*Category, error) {
	if name == "" {
		return nil, ErrCategoryNull
	}
	if !re.MatchString(name) {
		return nil, ErrCategoryInvalid
	}
	name = strings.ToLower(name)
	newCategory, err := NewCategory(userId, name, false)
	if err != nil {
		return nil, err
	}

	err = s.repo.AddCategory(newCategory)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (s *categoryService) GetCategoryById(categoryId, userId uuid.UUID) (*Category, error) {

	if userId.String() == "" {
		return nil, users.ErrUserIdNull
	}

	if categoryId.String() == "" {
		return nil, ErrCategoryIdNull
	}

	category, err := s.repo.GetCategoryByID(categoryId, userId)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) UpdateCategory(categoryId, userId uuid.UUID, name string) (*Category, error) {
	if userId.String() == "" {
		return nil, users.ErrUserIdNull
	}

	if categoryId.String() == "" {
		return nil, ErrCategoryIdNull
	}

	updatedCategory, err := s.repo.UpdateCategoryByID(categoryId, userId, name)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (s *categoryService) DeleteCategory(categoryId, userId uuid.UUID) error {
	if userId.String() == "" {
		return users.ErrUserIdNull
	}

	if categoryId.String() == "" {
		return ErrCategoryIdNull
	}

	err := s.repo.DeleteCategoryByID(categoryId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *categoryService) GetAllCategories(userId uuid.UUID) ([]Category, error) {
	var result []Category
	if userId.String() == "" {
		return result, users.ErrUserIdNull
	}

  result, err := s.repo.ListCategoriesByUser(userId)
  if err != nil {
    return result, err
  }

  return result, nil
}
