package category

import (
	"errors"
	"strings"
	"testing"
)

var (
	mockMap     = NewCategories()
	control     = "bills"
	invalidName = "bill!"
	emptyName   = ""
)

func deleteMap() {
	for k, _ := range mockMap.Store {
		delete(mockMap.Store, k)
	}
}


func TestAddCategory(t *testing.T) {
  err := mockMap.AddCategory(control)
  if err != nil {
    t.Error("Unexpected error occured: ", err)
  }
	if _, ok := mockMap.Store[strings.ToLower(control)]; !ok {
		t.Error(ErrCategoryNotAdded)
	}
	if control != mockMap.Store[strings.ToLower(control)].Name {
		t.Errorf("Expected category '%s' to be added but it was not found", control)
	}
}

func TestAddCategory_Duplicate(t *testing.T) {
  t.Cleanup(deleteMap)
  mockMap.AddCategory(control)
	category := "bills"
	err := mockMap.AddCategory(category)
	if !errors.Is(err, ErrCategoryExists) {
		t.Errorf("Expected error '%s' but got '%v'", ErrCategoryExists, err)
	}
}

func TestAddCategory_InvalidName(t *testing.T) {
  t.Cleanup(deleteMap)
	err := mockMap.AddCategory(invalidName)
	if !errors.Is(err, ErrCategoryInvalid) {
		t.Errorf("Expected error '%s' but got '%v' ", ErrCategoryInvalid, err)
	}
}

func TestAddCategory_EmptyName(t *testing.T) {
  t.Cleanup(deleteMap)
  err := mockMap.AddCategory(emptyName)
  if !errors.Is(err, ErrCategoryNull) {
    t.Errorf("Expected error '%s' but got '%v' ", ErrCategoryNull, err)
  }
}

func TestCreateDefaultCategories(t *testing.T) {
  t.Cleanup(deleteMap)
  err := mockMap.CreateDefaultCategories()
  if err != nil {
    t.Errorf("Test failed with error message '%v' while creating default categories. Map value: %v", err, mockMap.Store)
  }
  for _, category := range defaultCategories {
    if _, ok := mockMap.Store[strings.ToLower(category)]; !ok {
      t.Errorf("Category '%s' not found in map after CreateDefaultCategories", category)
    }
  }
}
