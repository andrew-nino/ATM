package service

import (
	"github.com/andrew-nino/ATM/config"
	"github.com/andrew-nino/ATM/entity"

	postgres "github.com/andrew-nino/ATM/internal/repository/postgresdb"
)

type Client interface {
	AddAccount(client entity.Client) (int, error)
}

type Service struct {
	Client
}

func NewService(reposPG *postgres.PG_Repository, cfg *config.Config) *Service {
	return &Service{
		Client:          NewClientService(reposPG.ClientPostgres, cfg),
	}
}
