package transaction

import (
	"testing"
	"time"
	"trackergo/internal/category"
	"trackergo/internal/users"
)

var (
	userMap        = users.NewInMemoryStore()
	transactionMap = NewInMemoryStore()
	categoryMap    = category.NewInMemoryStore()
	username       = "Testuser1234"
	email          = "testuser@test.com"
	password       = "MyStrongPassword123!"
	createdAt      = time.Now()
	amount         = 302.10
	newAmount      = 1000.0
)

func deleteMap() {
	for k, _ := range transactionMap.Store {
		delete(transactionMap.Store, k)
	}

	for k, _ := range userMap.Users {
		delete(userMap.Users, k)
	}

	for k, _ := range userMap.UserIDToEmail {
		delete(userMap.UserIDToEmail, k)
	}

	for k, _ := range userMap.UserIDToUsername {
		delete(userMap.UserIDToUsername, k)
	}

	for k, _ := range categoryMap.UserCategories {
		delete(categoryMap.UserCategories, k)
	}
}

func initialiseTest() (*users.User, *category.Category, error) {
	newUser := users.NewUser(username, email, password)

	testCategory, err := category.NewCategory("test", newUser.Id, false)
	if err != nil {
		return nil, nil, err
	}
	err = userMap.AddUser(newUser)
	if err != nil {
		return nil, nil, err
	}
	err = categoryMap.AddCategory(testCategory)
	if err != nil {
		return nil, nil, err
	}
	return newUser, testCategory, nil
}

func initialiseTransactionForTest() (*users.User, *category.Category, *Transaction, error) {
	testUser, testCategory, err := initialiseTest()
	if err != nil {
		return nil, nil, nil, err
	}
	newTransaction, err := NewTransaction(testCategory.Id, testUser.Id, amount, createdAt)

	err = transactionMap.AddTransaction(newTransaction)
	if err != nil {
		return nil, nil, nil, err
	}

	return testUser, testCategory, newTransaction, nil
}

func TestNewTransaction(t *testing.T) {
	t.Cleanup(deleteMap)
	testUser, testCategory, err := initialiseTest()
	if err != nil {
		t.Error(err)
	}
	_, err = NewTransaction(testUser.Id, testCategory.Id, amount, createdAt)
	if err != nil {
		t.Error(err)
	}
}

func TestAddTransactionInMemoryStore(t *testing.T) {
	t.Cleanup(deleteMap)
	testUser, testCategory, err := initialiseTest()
	if err != nil {
		t.Error(err)
	}

	newTransaction, err := NewTransaction(testUser.Id, testCategory.Id, amount, createdAt)
	if err != nil {
		t.Error(err)
	}

	err = transactionMap.AddTransaction(newTransaction)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTransactionFromInMemoryStore(t *testing.T) {
	t.Cleanup(deleteMap)
	testUser, _, testTransaction, err := initialiseTransactionForTest()
	if err != nil {
		t.Error(err)
	}
	_, err = transactionMap.GetTransaction(testTransaction.Id, testUser.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTransactionFromInMemoryStore(t *testing.T) {
	t.Cleanup(deleteMap)
	var newAmount float64
	testUser, _, testTransaction, err := initialiseTransactionForTest()
	if err != nil {
		t.Error(err)
	}
	newCategory, err := category.NewCategory("newTest", testUser.Id, false)
	if err != nil {
		t.Error(err)
	}
	updatedTransaction, err := transactionMap.UpdateTransaction(testTransaction.Id, testUser.Id, newCategory.Id, newAmount)
	if err != nil {
		t.Error(err)
	}
	if updatedTransaction.CategoryID != newCategory.Id {
		t.Errorf("Expected new category ID '%s' but got '%s'", newCategory.Id, updatedTransaction.CategoryID)
	}
	if updatedTransaction.Amount != newAmount {
		t.Errorf("Expected new amount '%v' but got '%v'", newAmount, updatedTransaction.Amount)
	}
}

func TestDeleteTransactionFromInMemoryStore(t *testing.T) {
	t.Cleanup(deleteMap)
	testUser, _, testTransaction, err := initialiseTransactionForTest()
	if err != nil {
		t.Error(err)
	}

	err = transactionMap.DeleteTransaction(testTransaction.Id, testUser.Id)
	if err != nil {
		t.Error(err)
	}

	userTransactions := transactionMap.Store[testUser.Id]
	transaction, ok := userTransactions[testTransaction.Id]
	if ok {
		t.Errorf("Was expecting transaction ID '%s' to not be found but it was found", transaction.Id)
	}
}

func TestListTransactionFromInMemoryStore(t *testing.T) {
	t.Cleanup(deleteMap)
	testUser, _, _, err := initialiseTransactionForTest()
	if err != nil {
		t.Error(err)
	}

  _, err = transactionMap.ListTransactions(testUser.Id)
  if err != nil {
    t.Error(err)
  }
}
