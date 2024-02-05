package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/olukkas/rinha-2024-golang/internal/entities"
	"github.com/olukkas/rinha-2024-golang/internal/services"
	"net/http"
	"strconv"
	"strings"
)

type WebTransactionServer struct {
	transactionService *services.TransactionService
	clientService      *services.ClientService
}

func NewWebTransactionServer(service *services.TransactionService) *WebTransactionServer {
	return &WebTransactionServer{transactionService: service}
}

func (wt *WebTransactionServer) Create(w http.ResponseWriter, r *http.Request) {
	idParam := strings.Split(r.URL.Path, "/")[2]
	clientId, _ := strconv.Atoi(idParam)

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
