package category

import (
	"errors"
	"testing"
)

var (
	mockMap     = NewInMemoryStore()
	control     = Category{}
	invalidName = Category{}
	emptyName   = Category{}
)

func deleteMap() {
	for k, _ := range mockMap.Store {
		delete(mockMap.Store, k)
	}
}

func TestAddCategory(t *testing.T) {
	c, err := NewCategory("bills")
	if err != nil {
		t.Error(err)
	}
	err = mockMap.AddCategory(c)
	if err != nil {
		t.Error("Unexpected error occured: ", err)
	}
	if _, ok := mockMap.Store[c.Id]; !ok {
		t.Error(ErrCategoryNotAdded)
	}
	if "bills" != mockMap.Store[c.Id].Name {
		t.Errorf("Expected category '%s' to be added but it was not found", control.Name)
	}
}

func TestAddCategory_InvalidName(t *testing.T) {
	t.Cleanup(deleteMap)

	_, err := NewCategory("bill!")
		if !errors.Is(err, ErrCategoryInvalid) {
			t.Errorf("Expected error '%s' but got '%v' ", ErrCategoryInvalid, err)
		}
	}

func TestNewCategory_EmptyName(t *testing.T) {
	t.Cleanup(deleteMap)

	_, err := NewCategory("")
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
}
