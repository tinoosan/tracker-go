package category

import (
	"errors"
	"strings"
	"testing"
)

var (
	mockMap     = make(map[string]*Category)
	control     = "bills"
	invalidName = "bill!"
	emptyName   = ""
)

func deleteMap() {
	for k, _ := range mockMap {
		delete(mockMap, k)
	}
}


func TestAddCategory(t *testing.T) {
  err := AddCategory(control, mockMap)
  if err != nil {
    t.Error("Unexpected error occured: ", err)
  }
	if _, ok := mockMap[strings.ToLower(control)]; !ok {
		t.Error(ErrCategoryNotAdded)
	}
	if control != mockMap[strings.ToLower(control)].Name {
		t.Errorf("Expected category '%s' to be added but it was not found", control)
	}
}

func TestAddCategory_Duplicate(t *testing.T) {
  t.Cleanup(deleteMap)
  AddCategory(control, mockMap)
	category := "bills"
	err := AddCategory(category, mockMap)
	if !errors.Is(err, ErrCategoryExists) {
		t.Errorf("Expected error '%s' but got '%v'", ErrCategoryExists, err)
	}
}

func TestAddCategory_InvalidName(t *testing.T) {
  t.Cleanup(deleteMap)
	err := AddCategory(invalidName, mockMap)
	if !errors.Is(err, ErrCategoryInvalid) {
		t.Errorf("Expected error '%s' but got '%v' ", ErrCategoryInvalid, err)
	}
}

func TestAddCategory_EmptyName(t *testing.T) {
  t.Cleanup(deleteMap)
  err := AddCategory(emptyName, mockMap)
  if !errors.Is(err, ErrCategoryNull) {
    t.Errorf("Expected error '%s' but got '%v' ", ErrCategoryNull, err)
  }
}

func TestCreateDefaultCategories(t *testing.T) {
  t.Cleanup(deleteMap)
  err := CreateDefaultCategories(mockMap)
  if err != nil {
    t.Errorf("Test failed with error message '%v' while creating default categories. Map value: %v", err, mockMap)
  }
  for _, category := range defaultCategories {
    if _, ok := mockMap[strings.ToLower(category)]; !ok {
      t.Errorf("Category '%s' not found in map after CreateDefaultCategories", category)
    }
  }
}
