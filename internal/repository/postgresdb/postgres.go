package postgresdb

import (
	"github.com/andrew-nino/ATM/entity"

	"github.com/jmoiron/sqlx"
)

type ClientPostgres interface {
	AddClient(client entity.Client) (int, error)
}

type AccountRepository interface {
	Deposit(entity.Transaction) error
	Withdraw(entity.Transaction) error
	GetBalance(int) float64
}

type PG_Repository struct {
	ClientPostgres
	AccountRepository
}

func NewPGRepository(db *sqlx.DB) *PG_Repository {
	return &PG_Repository{
		ClientPostgres:  NewClientToPostgres(db),
		AccountRepository: NewAccountRepository(db),
	}
}
