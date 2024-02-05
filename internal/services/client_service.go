package services

import (
	"errors"
	"github.com/olukkas/rinha-2024-golang/internal/entities"
	"github.com/olukkas/rinha-2024-golang/internal/repositories"
	"math"
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

func (c *ClientService) UpdateBalance(
	client *entities.Client,
	balance int,
	transactionType entities.TransactionType,
) error {
	var newBalance int

	if transactionType == entities.Debit {
		newBalance = client.Balance - balance

		if math.Abs(float64(newBalance)) > float64(client.Limit) {
			return NotEnoughBalanceErr
		}
	} else {
		newBalance = client.Balance + balance
	}

	err := c.clientRepo.UpdateBalance(client.ID, newBalance)
	if err != nil {
		return err
	}

	return nil
}
