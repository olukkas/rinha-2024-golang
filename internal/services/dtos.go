package services

import "time"

type CreateTransactionDto struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type TransactionDto struct {
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizado_em"`
}

type BalanceDto struct {
	Total int       `json:"total"`
	Limit int       `json:"limite"`
	Date  time.Time `json:"data_extrato"`
}

type GetStatementDto struct {
	Balance          BalanceDto       `json:"saldo"`
	LastTransactions []TransactionDto `json:"ultimas_transacoes"`
}
