package category

import (
	"fmt"
	"regexp"
	"strings"
	"trackergo/pkg/utils"

	"github.com/google/uuid"
)
type Categories struct {
  Store map[string]*Category
}

type Category struct {
	Id   uuid.UUID
	Name string
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

func NewCategories() *Categories {
  return &Categories{Store: make(map[string]*Category)}
}

func (c *Categories) CreateDefaultCategories() error {
	fmt.Println("Creating default categories...")
	for i := 0; i < len(defaultCategories); i++ {
		err := c.AddCategory(defaultCategories[i])
		if err != nil {
			fmt.Println(err)
			return err
		}
    cat := c.Store[defaultCategories[i]]
    cat.Default = true
	}
	fmt.Println("Default categories created successfuly")
	return nil
}

// This creates a category. This is for testing purposes to make it easy
// to create new instances without tying them to the AddCategory() method

func CreateCategory(name string) *Category {
  c := Category{
    Id: utils.GenerateUUID(),
    Name: name,
  }

  return &c
}

func (c *Categories) AddCategory(name string) error {

	if name == "" {
		return ErrCategoryNull
	}
	if !re.MatchString(name) {
		return ErrCategoryInvalid
	}
	name = strings.ToLower(name)
	_, ok := c.Store[name]
	if ok {
		return ErrCategoryExists
	}
 category := CreateCategory(name)
	c.Store[name] = category
	return nil
}
