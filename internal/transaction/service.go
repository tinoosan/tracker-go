package transaction

import (
	"time"
	"trackergo/internal/category"

	"github.com/google/uuid"
)

type TransactionService struct {
	repo            TransactionRepository
	categoryService *category.CategoryService
}

func NewTransactionService(repo TransactionRepository, categoryService *category.CategoryService) *TransactionService {
  return &TransactionService{repo: repo, categoryService: categoryService}
}

func (s *TransactionService) CreateTransaction(userId, categoryId uuid.UUID, amount float64, description string, createdAt time.Time) (*Transaction, error) {
	if createdAt.String() == "" {
		return nil, ErrDateNull
	}
	if categoryId.String() == "" {
		return nil, ErrTransactionCategoryNull
	}
	if amount == 0.0 {
		return nil, ErrAmountNull
	}
	if amount < 0.0 {
		return nil, ErrAmountNotPositive
	}

	newTransaction := NewTransaction(userId, categoryId, amount, description, createdAt)
	if newTransaction == nil {
		return nil, ErrTransactionNull
	}
	if err := s.repo.AddTransaction(newTransaction); err != nil {
		return nil, err
	}
	return newTransaction, nil
}

func (s *TransactionService) GetTransactionById(transactionId, userId uuid.UUID) (*Transaction, error) {
	if transactionId.String() == "" {
		return nil, ErrTransactionNull
	}
	transaction, err := s.repo.GetTransaction(transactionId, userId)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) UpdateTransaction(transactionId, userId, categoryId uuid.UUID, amount *float64) (*Transaction, error) {
	if transactionId.String() == "" || userId.String() == "" || categoryId.String() == "" {
		return nil, ErrTransactionNull
	}
	updateTransaction, err := s.repo.UpdateTransaction(transactionId, userId, categoryId, amount)
	if err != nil {
		return nil, err
	}
	return updateTransaction, nil
}

func (s *TransactionService) DeleteTransaction(transactionId, userId uuid.UUID) error {
	if transactionId.String() == "" || userId.String() == "" {
		return ErrTransactionNull
	}

	err := s.repo.DeleteTransaction(transactionId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionService) GetAllTransactions(userId uuid.UUID) ([]map[string]interface{}, error) {
	if userId.String() == "" {
		return nil, ErrTransactionNull
	}
	transactions, err := s.repo.ListTransactions(userId)
	if err != nil {
		return nil, err
	}

  var result []map[string]interface{}
  for _, t := range transactions {
    category, err := s.categoryService.GetCategoryById(t.CategoryID, userId)
    if err != nil {
      return nil, err
    }

    result = append(result, map[string]interface{}{
      "id": t.Id,
      "userId" : t.UserID,
      "categoryId": t.CategoryID,
      "categoryName": category.Name,
      "description": t.Description,
      "amount": t.Amount,
      "createdAt": t.CreatedAt,
    }) 
  }
	return result, nil
}
