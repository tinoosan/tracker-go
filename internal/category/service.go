package category

import (
	"strings"
	"trackergo/internal/users"

	"github.com/google/uuid"
)


type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateDefaultCategories() error {
	err := s.repo.CreateDefaultCategories()
	if err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) CreateCategory(userId uuid.UUID, name string) (*Category, error) {
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

func (s *CategoryService) GetCategoryById(categoryId, userId uuid.UUID) (*Category, error) {

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

func (s *CategoryService) UpdateCategory(categoryId, userId uuid.UUID, name string) (*Category, error) {
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

func (s *CategoryService) DeleteCategory(categoryId, userId uuid.UUID) error {
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

func (s *CategoryService) GetAllCategories(userId uuid.UUID) ([]Category, error) {
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
