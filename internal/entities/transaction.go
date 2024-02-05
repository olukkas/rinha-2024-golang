package entities

import (
	"errors"
	"time"
)

type TransactionType string

const (
	Credit = "c"
	Debit  = "d"
)

type Transaction struct {
	ID          int             `json:"id"`
	ClientID    int             `json:"client_id"`
	Value       int             `json:"valor"`
	Type        TransactionType `json:"tipo"`
	Description string          `json:"descrição"`
	CreatedAt   time.Time       `json:"realizada_em"`
}

func NewTransaction(value int, Type TransactionType, description string) (*Transaction, error) {
	transaction := Transaction{
		Value:       value,
		Type:        Type,
		Description: description,
		CreatedAt:   time.Now(),
	}

	if err := transaction.Validate(); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *Transaction) Validate() error {
	if t.Type == "" {
		return errors.New("o tipo da transação é requerido")
	}

	if t.Type != Credit && t.Type != Debit {
		return errors.New("tipo de transação não existe")
	}

	if t.Description == "" {
		return errors.New("descrição é requerida para a operação")
	}

	if len(t.Description) > 10 {
		return errors.New("o campo descrição deve conter entre 1 e 10 caracteres apenas")
	}

	if t.Value >= 0 {
		return errors.New("o valor para a transação deve ser maior do que 0")
	}

	return nil
}
