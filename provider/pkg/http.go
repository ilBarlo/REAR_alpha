package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// StartHttpServer starts a new HTTP server
func StartHttpServer() {
	// mux creation
	router := mux.NewRouter()

	// routes definition
	router.HandleFunc("/api/listflavours", listAllFlavoursHandler).Methods("GET")
	router.HandleFunc("/api/listflavours/selector", listAllFlavoursSelectorHandler).Methods("GET")
	router.HandleFunc("/api/reserveflavour/{flavourID}", reserveFlavourHandler).Methods("POST")
	router.HandleFunc("/api/purchaseflavour/{flavourID}", purchaseFlavourHandler).Methods("POST")
	router.HandleFunc("/api/listflavours/selector/syntax", getSyntaxes).Methods("GET")
	router.HandleFunc("/api/listflavours/selector/type", getTypes).Methods("GET")

	// Start server HTTP
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// listAllFlavoursHandler handles the request to get all Flavours
func listAllFlavoursHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve Flavours from MongoDB based on the architecture
	flavours, err := listAllFlavours()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavours)
}

func listAllFlavoursSelectorHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Parse the request body into the APIRequest struct
	var request Selector
	err = json.Unmarshal(body, &request)
	if err != nil {
		return
	}

	// Create the selector based on the parsed request
	selector := Selector{
		CPU:         request.CPU,
		RAM:         request.RAM,
		Policy:      request.Policy,
		FlavourType: request.FlavourType,
	}

	fmt.Println(selector)

	flmatch := findMatchingFlavours(selector)
	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flmatch)
}

func getSyntaxes(w http.ResponseWriter, r *http.Request) {
	selector := getSyntax()

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(selector)
}

func getTypes(w http.ResponseWriter, r *http.Request) {
	types := getFlavourTypes()

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types)
}

// Handler for reserving a Flavour
func reserveFlavourHandler(w http.ResponseWriter, r *http.Request) {
	// Get the flavourID value from the URL parameters
	params := mux.Vars(r)
	flavourID := params["flavourID"]

	// Check if the Flavour exists
	//
	//
	// REMEMBER

	// Create a new transaction
	transactionID := generateTransactionID()
	startTime := time.Now()
	transaction := Transaction{
		TransactionID: transactionID,
		FlavourID:     flavourID,
		StartTime:     startTime,
	}

	fmt.Println(transaction)

	// Initialize the transactions map if it is nil
	if transactions == nil {
		transactions = make(map[string]Transaction)
	}

	// Save the transaction in the transactions map
	transactions[transactionID] = transaction

	// // Start a 10-second timer
	// go func() {
	// 	<-time.After(10 * time.Second)
	// 	delete(transactions, transactionID)
	// }()

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// Handler for purchasing a Flavour
func purchaseFlavourHandler(w http.ResponseWriter, r *http.Request) {
	// Get the flavourID value from the URL parameters
	params := mux.Vars(r)
	flavourID := params["flavourID"]

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the request body into the APIRequest struct
	var purchase Purchase

	err = json.Unmarshal(body, &purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(purchase)

	// Check if the transactions map is nil or not initialized
	if transactions == nil {
		http.Error(w, "No active transactions found", http.StatusNotFound)
		return
	}

	// Retrieve the transaction from the transactions map
	transaction, exists := transactions[purchase.TransactionID]
	if !exists {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	// Check if the transaction is still valid (within 10 seconds)
	elapsedTime := time.Since(transaction.StartTime)
	if elapsedTime > 10*time.Second {
		http.Error(w, "Transaction Timeout", http.StatusInternalServerError)
		delete(transactions, purchase.TransactionID)
		return
	}

	// Perform the Flavour purchase
	// ...

	// Remove the transaction from the transactions map
	delete(transactions, purchase.TransactionID)

	// Respond with the FlavourID as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavourID)
}
