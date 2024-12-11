package category

import (
	"fmt"
	"regexp"
	"strings"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	AddCategory(c Category) error
	GetCategory(u uuid.UUID) (*Category, error)
	UpdateCategory(u uuid.UUID, name string) (*Category, error)
	DeleteCategory(u uuid.UUID) error
	ListCategories() []Category
}

type InMemoryStore struct {
	Store map[uuid.UUID]*Category
}

type Category struct {
	Id      uuid.UUID
	Name    string
	Default bool
}

type Error struct {
	message string
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
)

var (
	ErrCategoryExists   = &Error{message: "Category already exists"}
	ErrCategoryNull     = &Error{message: "Name cannot be empty"}
	ErrCategoryInvalid  = &Error{message: "Name contains invalid characters. Only characters a-z, A-Z and spaces are allowed"}
	ErrCategoryNotAdded = &Error{message: "Category could not be added"}
	ErrCategoryNotFound = &Error{message: "Category could not be found or does not exist"}
)

func (e *Error) Error() string {
	return e.message
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{Store: make(map[uuid.UUID]*Category)}
}

func (s *InMemoryStore) CreateDefaultCategories() error {
	fmt.Println("Creating default categories...")
	for i := 0; i < len(defaultCategories); i++ {
		c, err := NewCategory(defaultCategories[i])
		if err != nil {
			fmt.Println(err)
			return err
		}
    err = s.AddCategory(c)
	}
	fmt.Println("Default categories created successfuly")
	return nil
}

// This creates a category. This is for testing purposes to make it easy
// to create new instances without tying them to the AddCategory() method

func NewCategory(name string) (*Category, error) {
	if name == "" {
		return nil, ErrCategoryNull
	}
	if !re.MatchString(name) {
		return nil, ErrCategoryInvalid
	}
	name = strings.ToLower(name)
  c := &Category{}

	c.Id = utils.GenerateUUID()
  c.Name = name

	return c, nil
}

func (s *InMemoryStore) AddCategory(c *Category) error {

	_, ok := s.Store[c.Id]
	if ok {
		return ErrCategoryExists
	}
	s.Store[c.Id] = c
	return nil
}
