package main

import (
	"database/sql"
	"github.com/olukkas/rinha-2024-golang/internal/repositories"
	"github.com/olukkas/rinha-2024-golang/internal/server"
	"github.com/olukkas/rinha-2024-golang/internal/services"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DNS_DB"))
	if err != nil {
		log.Fatalf("unable to open db connection: %s", err)
	}
	defer db.Close()

	clientRepo := repositories.NewClientRepositoryDB(db)
	transactionRepo := repositories.NewTransactionRepositoryDB(db)

	clientService := services.NewClientService(clientRepo)
	transactionService := services.NewTransactionService(transactionRepo, clientService)

	web := server.NewWebTransactionServer(transactionService)

	http.HandleFunc("/clients/{id}/transacoes", web.Create)

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
