package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/olukkas/rinha-2024-golang/internal/entities"
	"github.com/olukkas/rinha-2024-golang/internal/services"
	"net/http"
	"strconv"
	"time"
)

type WebTransactionServer struct {
	transactionService *services.TransactionService
	clientService      *services.ClientService
}

func NewWebTransactionServer(
	transactionService *services.TransactionService,
	clientService *services.ClientService,
) *WebTransactionServer {
	return &WebTransactionServer{
		transactionService: transactionService,
		clientService:      clientService,
	}
}

func (wt *WebTransactionServer) Create(w http.ResponseWriter, r *http.Request) {
	paramId := chi.URLParam(r, "id")
	clientId, _ := strconv.Atoi(paramId)

	client, err := wt.clientService.GetById(clientId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var transaction entities.Transaction

	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction.CreatedAt = time.Now()

	if err = transaction.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result, err := wt.transactionService.CreateTransaction(client, &transaction)
	if err != nil {
		if errors.Is(err, services.NotEnoughBalanceErr) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondOk(w, result)
}

func respondOk(w http.ResponseWriter, result any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}
