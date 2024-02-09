package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/olukkas/rinha-2024-golang/internal/repositories"
	"github.com/olukkas/rinha-2024-golang/internal/server"
	"github.com/olukkas/rinha-2024-golang/internal/services"
	"log"
	"net/http"
	"os"
)

func main() {
	dns := os.Getenv("DNS_DB")
	if dns == "" {
		log.Fatalf("could not find DNS_DB env variable")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("could not find PORT env variable")
	}

	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatalf("unable to open db connection: %s", err)
	}
	defer db.Close()

	clientRepo := repositories.NewClientRepositoryDB(db)
	transactionRepo := repositories.NewTransactionRepositoryDB(db)

	clientService := services.NewClientService(clientRepo)
	transactionService := services.NewTransactionService(transactionRepo, clientService)

	web := server.NewWebTransactionServer(transactionService, clientService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Post("/clientes/{id}/transacoes", web.Create)
	c.Get("/clientes/{id}/extrato", web.Statement)

	fmt.Println("listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, c))
}
