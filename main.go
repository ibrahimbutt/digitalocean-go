package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func main() {
	// Seed the random number generator for generating dummy data
	rand.Seed(time.Now().UnixNano())

	// Create a slice of 10000 dummy transactions
	transactions := make([]Transaction, 10000)
	for i := 0; i < len(transactions); i++ {
		transactions[i] = Transaction{
			ID:          fmt.Sprintf("txn%d", i+1),
			Amount:      rand.Float64() * 100,
			Description: fmt.Sprintf("Transaction #%d", i+1),
			Date:        time.Now().Add(-time.Duration(rand.Intn(365)) * time.Hour * 24),
		}
	}

	// Set up a handler for the GET /transactions route
	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		// Convert the transactions slice to JSON
		transactionsJSON, err := json.Marshal(transactions)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the content type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the transactions JSON to the response
		w.Write(transactionsJSON)
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
