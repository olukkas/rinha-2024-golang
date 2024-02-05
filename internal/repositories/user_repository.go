package repositories

import (
	"database/sql"
	"github.com/olukkas/rinha-2024-golang/internal/entities"
)

type ClientRepository interface {
	UpdateBalance(id, balance int) error
	GetById(id int) (*entities.Client, error)
}

type ClientRepositoryDB struct {
	db *sql.DB
}

func NewClientRepositoryDB(db *sql.DB) *ClientRepositoryDB {
	return &ClientRepositoryDB{db: db}
}

//goland:noinspection SqlNoDataSourceInspection,SqlResolve
func (c *ClientRepositoryDB) GetById(id int) (*entities.Client, error) {
	var client entities.Client

	err := c.db.QueryRow("select id, balance, 'limit' where id = ?", id).
		Scan(&client.ID, &client.Balance, &client.Limit)

	if err != nil {
		return nil, err
	}

	return &client, nil
}

//goland:noinspection SqlNoDataSourceInspection,SqlResolve
func (c *ClientRepositoryDB) UpdateBalance(id, balance int) error {
	query := `update clients set balance = ? where id = ?`

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, balance)
	if err != nil {
		return err
	}

	return nil
}
