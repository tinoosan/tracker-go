package category

import (
	"testing"
	"trackergo/internal/users"
)

var (
	mockMap     = NewInMemoryStore()
	userMap     = users.NewInMemoryStore()
	username    = "Testuser1234"
	email       = "testuser@test.com"
	password    = "MyStrongPassword123!"
	invalidName = Category{}
	emptyName   = Category{}
)

func initialiseTest() (*users.User, *Category, error) {
	newUser, err := users.NewUser(username, email, password)
	if err != nil {
		return nil, nil, err
	}

	testCategory, err := NewCategory("test", newUser, false)
	if err != nil {
		return nil, nil, err
	}
	return newUser, testCategory, nil
}

func deleteMap() {
	for k, _ := range mockMap.DefaultCategories {
		delete(mockMap.DefaultCategories, k)
	}

	for k, _ := range mockMap.UserCategories {
		delete(mockMap.UserCategories, k)
	}
}

func TestAddDefaultCategory(t *testing.T) {
	t.Cleanup(deleteMap)
	_, testCategory, err := initialiseTest()
	if err != nil {
		t.Error(err)
	}
	err = mockMap.AddDefaultCategory(testCategory)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateDefaultCategory(t *testing.T) {
	t.Cleanup(deleteMap)
	err := mockMap.CreateDefaultCategories()
	if err != nil {
		t.Error(err)
	}
}

func TestAddCategory(t *testing.T) {
	t.Cleanup(deleteMap)
	_, testCategory, err := initialiseTest()
	if err != nil {
		t.Error(err)
	}
	err = mockMap.AddCategory(testCategory)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCategoryByID(t *testing.T) {
	t.Cleanup(deleteMap)
	testUser, testCategory, err := initialiseTest()
	if err != nil {
		t.Error(err)
	}

	err = mockMap.AddCategory(testCategory)
	if err != nil {
		t.Error(err)
	}

	_, err = mockMap.GetCategoryByID(testCategory.Id, testUser.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteCategoryByID(t *testing.T) {
	t.Cleanup(deleteMap)
	testUser, testCategory, err := initialiseTest()
	if err != nil {
		t.Error(err)
	}

	err = mockMap.AddCategory(testCategory)
	if err != nil {
		t.Error(err)
	}

  err = mockMap.DeleteCategoryByID(testCategory.Id, testUser.Id)
  if err != nil {
    t.Error(err)
  }
}
