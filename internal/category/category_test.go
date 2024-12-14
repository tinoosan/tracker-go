package category

import (
	"testing"
	"trackergo/internal/users"

	"github.com/google/uuid"
)

var (
	mockMap     = NewInMemoryStore()
  username = "testuser1234"
  email = "testuser@test.com"
  password = "testpassword"
	newUser, _     = users.NewUser(username, email, password)
	invalidName = Category{}
	bills, _    = NewCategory("bills", newUser, false)
	emptyName   = Category{}
)

func deleteMap() {
	for k, _ := range mockMap.DefaultCategories {
		delete(mockMap.DefaultCategories, k)
	}

	for k, _ := range mockMap.UserCategories {
		delete(mockMap.UserCategories, k)
	}
}

func TestAddCategory(t *testing.T) {
	userCategories := make(map[uuid.UUID]*Category)
	mockMap.UserCategories[newUser.Id] = userCategories

	err := mockMap.AddCategory(bills)
	if err != nil {
		t.Error("Unexpected error occured: ", err)
	}
	if _, ok := userCategories[bills.Id]; !ok {
		t.Error(ErrCategoryNotAdded)
	}
	if "bills" != userCategories[bills.Id].Name {
		t.Errorf("Expected category '%s' to be added but it was not found", bills.Name)
	}
}

func TestAddDefaultCategory(t *testing.T) {
  err := mockMap.AddDefaultCategory(bills)
  if err != nil {
    t.Error(err)
  }
}


func TestCreateDefaultCategory(t *testing.T) {
  err := mockMap.CreateDefaultCategories()
  if err != nil {
    t.Error(err)
  }
}



