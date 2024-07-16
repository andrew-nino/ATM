package entity

import "time"

type Client struct {
	ClientName string    `db:"client_name" json:"client_name" binding:"required"`
	Password   string    `db:"password" json:"password" binding:"required"`
	CreatedAt  time.Time `db:"created_at" json:"-"`
	UpdatedAt  time.Time `db:"update_at" json:"-"`
}

type Transaction struct {
	AccountId int `json:"-"`
	Amount float64 `json:"amount"`
}
