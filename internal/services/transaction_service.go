package services

import (
	"github.com/olukkas/rinha-2024-golang/internal/entities"
	"github.com/olukkas/rinha-2024-golang/internal/repositories"
	"math"
	"time"
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

func (t *TransactionService) CreateTransaction(
	client *entities.Client,
	transaction *entities.Transaction,
) (*CreateTransactionDto, error) {
	var newBalance int

	if transaction.Type == entities.Debit {
		newBalance = client.Balance - transaction.Value

		if math.Abs(float64(newBalance)) > float64(client.Balance) {
			return nil, NotEnoughBalanceErr
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

func (t *TransactionService) GetStatement(client *entities.Client) (*GetStatementDto, error) {
	var transactions []TransactionDto

	lastTransactions, err := t.transactionRepo.GetLastTransactions()
	if err != nil {
		return nil, err
	}

	for _, last := range lastTransactions {
		transaction := TransactionDto{
			Value:       last.Value,
			Type:        string(last.Type),
			Description: last.Description,
			CreatedAt:   last.CreatedAt,
		}

		transactions = append(transactions, transaction)
	}

	balance := BalanceDto{
		Total: client.Balance,
		Limit: client.Limit,
		Date:  time.Now(),
	}

	return &GetStatementDto{
		Balance:          balance,
		LastTransactions: transactions,
	}, nil
}
