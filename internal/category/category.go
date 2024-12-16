package category

import (
	"fmt"
	"regexp"
	"strings"
	"trackergo/internal/users"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	AddCategory(c *Category) error
	AddDefaultCategory(c *Category) error
	CreateDefaultCategories() error
	GetCategoryByID(categoryId, userId uuid.UUID) (*Category, error)
	UpdateCategoryByID(categoryId, userId uuid.UUID, name string) (*Category, error)
	DeleteCategoryByID(categoryId, userId uuid.UUID) error
  ListCategoriesByUser(userId uuid.UUID) ([]Category, error)
}

type InMemoryStore struct {
	DefaultCategories map[uuid.UUID]*Category
	UserCategories    map[uuid.UUID]map[uuid.UUID]*Category
}

type Category struct {
	Id      uuid.UUID
	Name    string
	User    *users.User
	Default bool
}

var (
	pattern           = `^[a-zA-Z\s]+$`
	re                = regexp.MustCompile(pattern)
	defaultCategories = [4]string{
		"bills",
		"rent",
		"transport",
		"entertainment",
	}
	_ CategoryRepository = &InMemoryStore{}
)

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{DefaultCategories: make(map[uuid.UUID]*Category),
		UserCategories: make(map[uuid.UUID]map[uuid.UUID]*Category)}
}

func NewCategory(name string, user *users.User, isDefault bool) (*Category, error) {
	if name == "" {
		return nil, ErrCategoryNull
	}
	if !re.MatchString(name) {
		return nil, ErrCategoryInvalid
	}
	name = strings.ToLower(name)
	c := &Category{
		Id:      utils.GenerateUUID(),
		Name:    name,
		User:    user,
		Default: isDefault,
	}

	return c, nil
}

func (s *InMemoryStore) CreateDefaultCategories() error {
	fmt.Println("Creating default categories...")
	for i := 0; i < len(defaultCategories); i++ {
		c, err := NewCategory(defaultCategories[i], users.SystemUser, true)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = s.AddDefaultCategory(c)
	}
	fmt.Println("Default categories created successfuly")
	return nil
}

func (s *InMemoryStore) AddDefaultCategory(category *Category) error {
	for _, v := range s.DefaultCategories {
		if v.Name == category.Name {
			return ErrCategoryExists
		}
	}
	s.DefaultCategories[category.Id] = category
	return nil
}

func (s *InMemoryStore) AddCategory(category *Category) error {
	if category.User == nil {
		return ErrCategoryHasNoUser
	}
	userId := category.User.Id
	userCategories, ok := s.UserCategories[userId]
	if !ok {
		userCategories = make(map[uuid.UUID]*Category)
		s.UserCategories[userId] = userCategories
	}

	for _, c := range s.UserCategories[userId] {
		if c.Name == category.Name {
			return ErrCategoryExists
		}
	}

	userCategories[category.Id] = category
	return nil
}

func (s *InMemoryStore) GetCategoryByID(categoryId, userId uuid.UUID) (*Category, error) {
	if userId.String() == "" {
		return nil, users.ErrUserIdNull
	}
	userCategories, ok := s.UserCategories[userId]
	if !ok {
		return nil, ErrUserHasNoCategories
	}

	if categoryId.String() == "" {
		return nil, ErrCategoryIdNull
	}

	category := userCategories[categoryId]
	return category, nil
}

func (s InMemoryStore) UpdateCategoryByID(categoryId, userId uuid.UUID, name string) (*Category, error) {
	if userId.String() == "" {
		return nil, users.ErrUserIdNull
	}
	userCategories, ok := s.UserCategories[userId]
	if !ok {
		return nil, ErrUserHasNoCategories
	}

	if categoryId.String() == "" {
		return nil, ErrCategoryIdNull
	}

  category, ok := userCategories[categoryId]
  if !ok {
    return nil, ErrCategoryNotFound
  }
  if name != "" && name != category.Name {
    category.Name = name
  }
  return category, nil
}

func (s *InMemoryStore) DeleteCategoryByID(categoryId uuid.UUID, userId uuid.UUID) error {
	if userId.String() == "" {
		return users.ErrUserIdNull
	}
	userCategories, ok := s.UserCategories[userId]
	if !ok {
		return ErrUserHasNoCategories
	}

	if categoryId.String() == "" {
		return ErrCategoryIdNull
	}

	delete(userCategories, categoryId)
	return nil

}

func (s *InMemoryStore) ListCategoriesByUser(userId uuid.UUID) ([]Category, error) {
  var result []Category
  if userId.String() == "" {
    return result, users.ErrUserIdNull
  }

  for _, defaultCategories := range s.DefaultCategories {
    result = append(result, *defaultCategories)
  }

  userCategories, ok := s.UserCategories[userId]
  if !ok {
    return result, ErrUserHasNoCategories
  }

  for _, category := range userCategories {
    result = append(result, *category)
  }
  return result, nil
}

