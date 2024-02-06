package services

import (
	"errors"
	"github.com/olukkas/rinha-2024-golang/internal/entities"
	"github.com/olukkas/rinha-2024-golang/internal/repositories"
)

var (
	NotEnoughBalanceErr = errors.New("saldo insuficiente")
)

type ClientService struct {
	clientRepo repositories.ClientRepository
}

func NewClientService(userRepo repositories.ClientRepository) *ClientService {
	return &ClientService{clientRepo: userRepo}
}

func (c *ClientService) GetById(id int) (*entities.Client, error) {
	return c.clientRepo.GetById(id)
}
