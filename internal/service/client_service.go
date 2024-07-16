package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/andrew-nino/ATM/config"
	"github.com/andrew-nino/ATM/entity"
	"github.com/andrew-nino/ATM/internal/repository/postgresdb"

	log "github.com/sirupsen/logrus"
)

type ClientService struct {
	repo   postgresdb.ClientPostgres
	config config.Config
}

func NewClientService(repo postgresdb.ClientPostgres, cfg *config.Config) *ClientService {
	return &ClientService{
		repo:   repo,
		config: *cfg}
}

func (s *ClientService) AddAccount(client entity.Client) (int, error) {

	password, err := generatePasswordHash(client.Password, s.config)
	if err != nil {
		return 0, err
	}
	client.Password = password
	return s.repo.AddAccount(client)
}

// generatePasswordHash generates a SHA1 hash of the given password with a salt.
// The salt is a constant string.
func generatePasswordHash(password string, cfg config.Config) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		log.Debugf("failed to generate password hash: %s", err)
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(cfg.SHA.Salt))), nil
}
