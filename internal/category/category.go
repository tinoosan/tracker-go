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
		c, err := NewCategory(defaultCategories[i], nil, true)
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

