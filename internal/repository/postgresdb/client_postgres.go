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

// Adding client information to the database.
func (c *ClientToPostgres) AddAccount(add entity.Client) (int, error) {
	var clientID int
	queryToClient := fmt.Sprintf(`INSERT INTO %s (client_name, password_hash) values ($1, $2) RETURNING id`, clientTable)
	rowClient := c.db.QueryRow(queryToClient, add.ClientName, add.Password)
	err := rowClient.Scan(&clientID)
	if err != nil {
		log.Debugf("repository.AddClient - rowClient.Scan : %v", err)
		return 0, err
	}

	return clientID, nil
}
