package application

import (
	"trackergo/internal/domain/ledger"
  vo "trackergo/internal/domain/valueobjects"

	"github.com/google/uuid"
)

type AccountGetter interface {
	FindByCode(code vo.Code, userID uuid.UUID) (*ledger.Account, error)
	FindByName(userID uuid.UUID, name string) (*ledger.Account, error)
	List(userID uuid.UUID) ([]*ledger.Account, error)
}

type AccountRepository interface {
	AccountGetter
	Save(account *ledger.Account) error
	Update(code vo.Code, userID uuid.UUID, name string) error
	Delete(code vo.Code, userID uuid.UUID) error
}

type LedgerRepository interface {
	Save(transaction *ledger.Entry) error
	FindByID(transactionId, userId uuid.UUID) (*ledger.Entry, error)
	Update(transactionId, userId uuid.UUID, amount *float64) error
	Delete(transactionId, userId uuid.UUID) error
	List(userId uuid.UUID) ([]*ledger.Entry, error)
}
