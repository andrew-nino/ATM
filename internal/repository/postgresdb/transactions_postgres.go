package postgresdb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andrew-nino/ATM/entity"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (a *AccountPostgres) Deposit(data entity.Transaction) error {
	ctx := context.Background()
	tx, err := a.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET balance = balance + $1, update_at = now() WHERE client_id = $2", accountsTable)
	_, err = tx.Exec(query, data.Amount, data.AccountId)
	if err != nil {
		log.Debugf("repository.Deposit - Exec : %v", err)
		tx.Rollback()
		return err
	}
	log.Infof("The top-up operation for account No. %d  was successful.", data.AccountId)
	return tx.Commit()
}

func (a *AccountPostgres) Withdraw(data entity.Transaction) error {
	ctx := context.Background()
	tx, err := a.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET balance = balance - $1, update_at = now() WHERE client_id = $2", accountsTable)
	_, err = tx.Exec(query, data.Amount, data.AccountId)
	if err != nil {
		log.Debugf("repository.Withdraw - Exec : %v", err)
		tx.Rollback()
		return err
	}
	log.Infof("The operation to withdraw from account No. %d was successful.", data.AccountId)
	return tx.Commit()
}

func (a *AccountPostgres) GetBalance(client_id int) float64 {
	var balance float64
	query := fmt.Sprintf("SELECT balance FROM %s WHERE client_id = $1", accountsTable)
	row := a.db.QueryRow(query, client_id)
	err := row.Scan(&balance)
	if err != nil {
		log.Debugf("repository.GetBalance - row.Scan : %v", err)
		return 0
	}
	log.Infof("The operation get balance from account No. %d was successful.", client_id)
	return balance
}
