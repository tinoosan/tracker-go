package users

import "testing"

var (
	username = "Testuser1234"
	email    = "testuser@test.com"
	password = "MyStrongPassword123!"
	mockMap  = NewInMemoryStore()
)

func deleteMap() {
	for k, _ := range mockMap.Users {
		delete(mockMap.Users, k)
	}

	for k, _ := range mockMap.UserIDToUsername {
		delete(mockMap.UserIDToUsername, k)
	}

	for k, _ := range mockMap.UserIDToEmail {
		delete(mockMap.UserIDToEmail, k)
	}
}

func TestNewUser(t *testing.T) {
	t.Cleanup(deleteMap)
	user := NewUser(username, email, password)

	if user == nil {
		t.Error(ErrUserNotCreated)
	}

}

func TestAddUser(t *testing.T) {
	t.Cleanup(deleteMap)
	user := NewUser(username, email, password)
	err := mockMap.AddUser(user)
	if err != nil {
		t.Error(err)
	}
}

func TestGetUserByID(t *testing.T) {
	t.Cleanup(deleteMap)
	user := NewUser(username, email, password)
	err := mockMap.AddUser(user)
	if err != nil {
		t.Error(err)
	}
	fetchedUser, _ := mockMap.GetUserByID(user.Id)
	if fetchedUser.Id != user.Id {
		t.Errorf("Expected user with ID '%s' but got '%s'", user.Id, fetchedUser.Id)
	}
}

func TestUpdateUserByID(t *testing.T) {
	t.Cleanup(deleteMap)
	user := NewUser(username, email, password)
	err := mockMap.AddUser(user)
	if err != nil {
		t.Error(err)
	}
	_, err = mockMap.UpdateUserByID(user.Id, "newuser", "newemail@test.com")
	if err != nil {
		t.Error(err)
	}

}

func TestDeleteUserByID(t *testing.T) {
	t.Cleanup(deleteMap)
	user := NewUser(username, email, password)
  err := mockMap.AddUser(user)
  if err != nil {
    t.Error(err)
  }
  err = mockMap.DeleteUserByID(user.Id)
  if err != nil {
    t.Error(err)
  }
}
