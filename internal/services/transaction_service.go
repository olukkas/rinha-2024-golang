package services

import (
	"github.com/olukkas/rinha-2024-golang/internal/entities"
	"github.com/olukkas/rinha-2024-golang/internal/repositories"
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
	_, err := t.transactionRepo.Save(transaction)
	if err != nil {
		return nil, err
	}

	err = t.clientService.UpdateBalance(client, transaction.Value, transaction.Type)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionDto{
		Limit:   client.Limit,
		Balance: client.Balance,
	}, nil
}
