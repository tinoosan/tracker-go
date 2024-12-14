package users

import "testing"

var (
  username = "Testuser1234"
  email = "testuser@test.com"
  password = "MyStrongPassword123!"
  mockMap = NewInMemoryStore()
)

func deleteMap() {
	for k, _ := range mockMap.Users{
		delete(mockMap.Users, k)
	}

	for k, _ := range mockMap.UserIDToUsername{
		delete(mockMap.UserIDToUsername, k)
	}
}

func TestNewUser(t *testing.T) {
  user, err := NewUser(username, email, password)
  if err != nil {
    t.Error(err)
  }

  if user == nil {
    t.Error(ErrUserNotCreated)
  }

}

func TestAddUser(t *testing.T) {
  t.Cleanup(deleteMap)
  user, _ := NewUser(username, email, password)
  err := mockMap.AddUser(user)
  if err != nil {
    t.Error(err)
  }
}
