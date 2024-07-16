package postgresdb

import (
	"github.com/andrew-nino/ATM/entity"

	"github.com/jmoiron/sqlx"
)

type ClientPostgres interface {
	AddAccount(client entity.Client) (int, error)
}

type PG_Repository struct {
	ClientPostgres
}

func NewPGRepository(db *sqlx.DB) *PG_Repository {
	return &PG_Repository{
		ClientPostgres:          NewClientToPostgres(db),
	}
}
