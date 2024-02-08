package repositories

import (
	"database/sql"
	"github.com/olukkas/rinha-2024-golang/internal/entities"
)

type ClientRepository interface {
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

	err := c.db.QueryRow("select id, balance, total_limit from clients where id = $1", id).
		Scan(&client.ID, &client.Balance, &client.Limit)

	if err != nil {
		return nil, err
	}

	return &client, nil
}
