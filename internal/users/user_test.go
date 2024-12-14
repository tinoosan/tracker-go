package users

import "testing"

var (
  username = "Testuser1234"
  email = "testuser@test.com"
  password = "MyStrongPassword123!"
)
func TestNewUser(t *testing.T) {
  user, err := NewUser(username, email, password)
  if err != nil {
    t.Error(err)
  }

  if user == nil {
    t.Error(ErrUserNotCreated)
  }

}
