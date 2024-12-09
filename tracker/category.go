package tracker

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type Category struct {
	Id   uuid.UUID
	Name string
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

func (c *Category) generateUUID() {
  c.Id = uuid.New()
}

func (e *Error) Error() string {
	return e.message
}

func CreateDefaultCategories(c map[string]*Category) error {
	fmt.Println("Creating default categories...")
	for i := 0; i < len(defaultCategories); i++ {
		err := AddCategory(defaultCategories[i], c)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	fmt.Println("Default categories created successfuly")
	return nil
}

func AddCategory(name string, c map[string]*Category) error {

	if name == "" {
		return ErrCategoryNull
	}
	if !re.MatchString(name) {
		return ErrCategoryInvalid
	}
	name = strings.ToLower(name)
	_, ok := c[name]
	if ok {
		return ErrCategoryExists
	}
  category := &Category{Name: name}
  category.generateUUID()
	c[name] = category
	return nil
}
