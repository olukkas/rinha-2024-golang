package repositories

import (
	"database/sql"
	"github.com/olukkas/rinha-2024-golang/internal/entities"
)

type TransactionRepository interface {
	Save(transaction *entities.Transaction, client *entities.Client) (*entities.Transaction, error)
	GetLastTransactions() ([]*entities.Transaction, error)
}

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{db: db}
}

//goland:noinspection SqlNoDataSourceInspection,SqlResolve
func (t *TransactionRepositoryDB) Save(
	tr *entities.Transaction,
	client *entities.Client,
) (*entities.Transaction, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	insertQuery := `insert into transactions (client_id, value, type, description, created_at)
	values ($1, $2, $3, $4, $5) `

	_, err = tx.Exec(insertQuery, tr.ClientID, tr.Value, tr.Type, tr.Description, tr.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	updateClientBalance := `update clients set balance = $1 where id = $2`

	_, err = tx.Exec(updateClientBalance, client.Balance, client.ID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return tr, nil
}

//goland:noinspection SqlNoDataSourceInspection,SqlResolve
func (t *TransactionRepositoryDB) GetLastTransactions() ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction

	query := `
	select id, client_id, value, type, description, created_at
	from transactions
	order by created_at
	limit 10
	`

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tr := new(entities.Transaction)

		err := rows.Scan(&tr.ID, &tr.ClientID, &tr.Value, &tr.Description, &tr.CreatedAt)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tr)
	}

	return transactions, nil
}
