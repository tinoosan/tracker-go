package tracker

import (
	"fmt"
	"regexp"
	"strings"
)

type Category struct {
	Name string
}

type Error struct {
	message string
}

var (
	c                 = make(map[string]*Category)
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
	ErrCategoryExists  = &Error{message: "Category already exists"}
	ErrCategoryNull    = &Error{message: "Name cannot be empty"}
	ErrCategoryInvalid = &Error{message: "Name contains invalid characters. Only characters a-z, A-Z and spaces are allowed"}
)

func (e *Error) Error() string {
	return e.message
}

func CreateDefaultCategories() {
	fmt.Println("Creating default categories...")
	for i := 0; i < len(defaultCategories); i++ {
		err := AddCategory(defaultCategories[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Default categories created successfuly")
}

func AddCategory(name string) error {

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
	c[name] = &Category{Name: name}
	return nil
}
