package service

import (
	"github.com/andrew-nino/ATM/config"
	"github.com/andrew-nino/ATM/entity"

	postgres "github.com/andrew-nino/ATM/internal/repository/postgresdb"
)

type Client interface {
	AddClient(client entity.Client) (int, error)
}

type BankAccount interface {
	Deposit(entity.Transaction) error
	Withdraw(entity.Transaction) error
	GetBalance(int) float64
}

type Service struct {
	Client
	BankAccount
}

func NewService(reposPG *postgres.PG_Repository, cfg *config.Config) *Service {
	return &Service{
		Client:      NewClientService(reposPG.ClientPostgres, cfg),
		BankAccount: NewAccountService(reposPG),
	}
}
