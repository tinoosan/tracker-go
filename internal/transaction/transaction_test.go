package transaction

import (
	"fmt"
	"testing"
	"time"
	"trackergo/internal/category"
	"trackergo/internal/users"
)

var (
	transactions = NewInMemoryStore()
	newUser      = users.NewUser()
	bills, _     = category.NewCategory("bills", newUser, false)
	rent, _      = category.NewCategory("rent", newUser, false)
	createdAt    = time.Now()
	amount       = 302.10
	newAmount    = 1000.0
	transaction  = Transaction{}
)

func deleteMap() {
	for k, _ := range transactions.Store {
		delete(transactions.Store, k)
	}
}

func TestCreateTransaction(t *testing.T) {
	transaction.NewTransaction(createdAt, bills, amount)
	err := transactions.AddTransaction(transaction)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTransactionFromInMemoryStore(t *testing.T) {
	t.Cleanup(deleteMap)
	transaction.NewTransaction(createdAt, bills, amount)
	err := transactions.AddTransaction(transaction)
	if err != nil {
		t.Error(err)
		return
	}

	testTransaction, err := transactions.GetTransaction(transaction.Id)
	if err != nil {
		t.Error(err)
		return
	}

	if testTransaction.Id != transaction.Id {
		t.Errorf("Expected transaction with Id '%s' but got '%s'", transaction.Id, testTransaction.Id)
	}
	fmt.Printf("Transaction %+v has been retrieved successfully", testTransaction)
}

func TestUpdateTransactionInMemoryStore(t *testing.T) {
	t.Cleanup(deleteMap)
	transaction.NewTransaction(createdAt, bills, amount)
	err := transactions.AddTransaction(transaction)
	if err != nil {
		t.Error(err)
		return
	}

	updatedTransaction, err := transactions.UpdateTransaction(transaction.Id, *rent, newAmount)
	if err != nil {
		t.Error(err)
		return
	}

	if updatedTransaction.Category.Name != rent.Name {
		t.Errorf("Expected transaction with updated name '%s' but got '%s'", rent.Name, updatedTransaction.Category.Name)
		return
	}

	if updatedTransaction.Amount != newAmount {
		t.Errorf("Expected transaction with updated amount '%v' but got '%v'", newAmount, updatedTransaction.Amount)
		return
	}
	fmt.Printf("Transaction %+v has been updated successfully", updatedTransaction)
}

func TestDeleteTransactionFromInMemoryStore(t *testing.T) {
	t.Cleanup(deleteMap)
	transaction.NewTransaction(createdAt, bills, amount)
	err := transactions.AddTransaction(transaction)
	if err != nil {
		t.Error(err)
		return
	}

	err = transactions.DeleteTransaction(transaction.Id)
	if err != nil {
		t.Error(err)
		return
	}

	if _, ok := transactions.Store[transaction.Id]; ok {
		t.Errorf("Expected transaction with Id '%s' to be deleted but it was found", transaction.Id)
		return
	}
	fmt.Printf("Transaction %+v has been deleted successfully", transaction.Id)
}
