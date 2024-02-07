package services

import (
	"errors"
	"github.com/olukkas/rinha-2024-golang/internal/entities"
	"github.com/olukkas/rinha-2024-golang/internal/repositories"
	"math"
)

type TransactionService struct {
	transactionRepo repositories.TransactionRepository
	clientService   *ClientService
}

func NewTransactionService(
	transactionRepo repositories.TransactionRepository,
	clientService *ClientService,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		clientService:   clientService,
	}
}

type CreateTransactionDto struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

func (t *TransactionService) CreateTransaction(
	client *entities.Client,
	transaction *entities.Transaction,
) (*CreateTransactionDto, error) {
	var newBalance int

	if transaction.Type == entities.Debit {
		newBalance = client.Balance - transaction.Value

		if math.Abs(float64(newBalance)) > float64(client.Balance) {
			return nil, errors.New("")
		}
	} else {
		newBalance = client.Balance + transaction.Value
	}

	client.Balance = newBalance

	_, err := t.transactionRepo.Save(transaction, client)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionDto{
		Limit:   client.Limit,
		Balance: client.Balance,
	}, nil
}
