package service

import (
	"github.com/andrew-nino/ATM/entity"
	"github.com/andrew-nino/ATM/internal/repository/postgresdb"
)

type AccountService struct {
	repo postgresdb.AccountRepository
}

func NewAccountService(repo postgresdb.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (a *AccountService) Deposit(data entity.Transaction) error {
	return a.repo.Deposit(data)
}

func (a *AccountService) Withdraw(data entity.Transaction) error {
	return a.repo.Withdraw(data)
}

func (a *AccountService) GetBalance(client_id int) float64 {
	return a.repo.GetBalance(client_id)
}
