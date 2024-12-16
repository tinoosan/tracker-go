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

	testCategory, err := NewCategory("test", newUser.Id, false)
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

func TestUpdateCategoryByID(t *testing.T) {
	t.Cleanup(deleteMap)
	testUser, testCategory, err := initialiseTest()
	if err != nil {
		t.Error(err)
	}
	err = mockMap.AddCategory(testCategory)
	newName := "newName"
	updatedCategory, err := mockMap.UpdateCategoryByID(testCategory.Id, testUser.Id, newName)
	if err != nil {
		t.Error(err)
	}
	if updatedCategory.Name != newName {
		t.Errorf("Expected update category with name '%s' but got '%s'", newName, updatedCategory.Name)
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

func TestListCategoriesByUser(t *testing.T) {
	t.Cleanup(deleteMap)
	err := mockMap.CreateDefaultCategories()
	if err != nil {
		t.Error(err)
	}

	testUser, testCategory, err := initialiseTest()
	if err != nil {
		t.Error(err)
	}
	err = mockMap.AddCategory(testCategory)
	if err != nil {
		t.Error(err)
	}

  list, err := mockMap.ListCategoriesByUser(testUser.Id)
  if err != nil {
    t.Error(err)
  }
  
  if len(list) == 0 {
    t.Error("List is empty")
  }

}
