package tracker

import "testing"

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
	category := "bills"
	AddCategory(category, mockMap)
	_, ok := mockMap[category]
	if !ok {
		t.Error(ErrCategoryNotAdded)
	}
	if category != mockMap[category].Name {
		t.Errorf("Expected category '%s' to be added but it was not found", category)
	}
}

func TestAddCategory_Duplicate(t *testing.T) {
	category := "bills"
	err := AddCategory(category, mockMap)
	if err != ErrCategoryExists {
		t.Errorf("Expected error '%s' but got '%v'", ErrCategoryExists, err)
	}

}

func TestAddCategory_InvalidName(t *testing.T) {
	err := AddCategory(invalidName, mockMap)
	if err != ErrCategoryInvalid {
		t.Errorf("Expected error '%s' but got '%v' ", ErrCategoryInvalid, err)
	}
}

func TestAddCategory_EmptyName(t *testing.T) {
  err := AddCategory(emptyName, mockMap)
  if err != ErrCategoryNull {
    t.Errorf("Expected error '%s' but got '%v' ", ErrCategoryNull, err)
  }
  t.Cleanup(deleteMap)
}

func TestCreateDefaultCategories(t *testing.T) {
  err := CreateDefaultCategories(mockMap)
  if err != nil {
    t.Errorf("Test failed with error message '%v ", err)
  }
}
