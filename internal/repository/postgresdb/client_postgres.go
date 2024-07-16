package postgresdb

import (
	"fmt"

	"github.com/andrew-nino/ATM/entity"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type ClientToPostgres struct {
	db *sqlx.DB
}

func NewClientToPostgres(db *sqlx.DB) *ClientToPostgres {
	return &ClientToPostgres{db: db}
}

// Adding customer information to the database is wrapped in a transaction.
// At the same time, we fill the client_id of the accaunts table with the received id.
func (c *ClientToPostgres) AddClient(add entity.Client) (int, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return 0, err
	}
	var clientID int
	queryToClient := fmt.Sprintf(`INSERT INTO %s (client_name, password_hash) values ($1, $2) RETURNING id`, clientTable)
	rowClient := tx.QueryRow(queryToClient, add.ClientName, add.Password)
	err = rowClient.Scan(&clientID)
	if err != nil {
		log.Debugf("repository.AddClient - rowClient.Scan : %v", err)
		tx.Rollback()
		return 0, err
	}
	queryToStatus := fmt.Sprintf(`INSERT INTO %s (client_id) values ($1) RETURNING id`, accountsTable)
	_, err = tx.Exec(queryToStatus, clientID)
	if err != nil {
		tx.Rollback()
		log.Debugf("repository.AddClient - tx.Exec : %v", err)
		return 0, err
	}
	log.Infof("The client adding operation %s fwas successful.", add.ClientName)
	return clientID, tx.Commit()
}
