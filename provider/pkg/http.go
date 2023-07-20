package pkg

import (
	"encoding/json"
	"fmt"
	"io"
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
	router.HandleFunc("/api/listflavours/{flavourID}", getFlavourByIDHandler).Methods("GET")
	router.HandleFunc("/api/listflavours/selector", listAllFlavoursSelectorHandler).Methods("POST")
	router.HandleFunc("/api/reserveflavour/{flavourID}", reserveFlavourHandler).Methods("POST")
	router.HandleFunc("/api/purchaseflavour/{flavourID}", purchaseFlavourHandler).Methods("POST")
	router.HandleFunc("/api/listflavours/selector/syntax", getSyntaxes).Methods("GET")
	router.HandleFunc("/api/listflavours/selector/type", getTypes).Methods("GET")

	// Start server HTTP
	fmt.Println("Server listening on port 14144")
	log.Fatal(http.ListenAndServe(":14144", router))
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

// listAllFlavoursSelectorHandler list all flavours that fit the selector
func listAllFlavoursSelectorHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
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
		FlavourType: request.FlavourType,
		CPU:         request.CPU,
		RAM:         request.RAM,
		MoreThanCPU: request.MoreThanCPU,
		MoreThanRAM: request.MoreThanRAM,
		LessThanCPU: request.LessThanCPU,
		LessThanRAM: request.LessThanRAM,
	}

	if selector.MoreThanCPU != 0 || selector.MoreThanRAM != 0 {
		flmatch := findMatchingFlavoursMore(selector)
		// Respond with the Flavours as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flmatch)
	} else if selector.LessThanCPU != 0 || selector.LessThanRAM != 0 {
		flmatch := findMatchingFlavoursLess(selector)
		// Respond with the Flavours as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flmatch)
	} else {
		flmatch := findMatchingFlavours(selector)
		// Respond with the Flavours as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(flmatch)
	}

}

// getFlavourByIDHandler gets the single Flavour by its FlavourID
func getFlavourByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the flavourID value from the URL parameters
	params := mux.Vars(r)
	flavourID := params["flavourID"]

	flavour, err := getFlavourByID(flavourID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavour)
}

// getSyntaxes gets the selector syntax
func getSyntaxes(w http.ResponseWriter, r *http.Request) {
	selector := getSyntax()

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(selector)
}

// getTypes get all the available Flavours' types
func getTypes(w http.ResponseWriter, r *http.Request) {
	typesFlavour := getFlavourTypes()
	typeList := []Type{}

	for _, typeF := range typesFlavour {
		for _, typeS := range types {
			if typeF == typeS.FlavourType {
				typeList = append(typeList, typeS)
			}
		}
	}
	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(typeList)
}

// reserveFlavourHandler reserves a Flavour by its flavourID
func reserveFlavourHandler(w http.ResponseWriter, r *http.Request) {
	// Get the flavourID value from the URL parameters
	params := mux.Vars(r)
	flavourID := params["flavourID"]

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var request struct {
		FlavourID string `json:"flavourID"`
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if flavourID != request.FlavourID {
		http.Error(w, "Mismatch body & param", http.StatusConflict)
		return
	}

	flavour, _ := getFlavourByID(flavourID)
	if flavour == nil {
		http.Error(w, "Flavour not found", http.StatusNotFound)
		return
	}

	// Create a new transaction
	transactionID := generateTransactionID()
	startTime := time.Now()
	transaction := Transaction{
		TransactionID: transactionID,
		FlavourID:     flavourID,
		StartTime:     startTime,
	}

	// Initialize the transactions map if it is nil
	if transactions == nil {
		transactions = make(map[string]Transaction)
	}

	// Save the transaction in the transactions map
	transactions[transactionID] = transaction

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// purchaseFlavourHandler is an handler for purchasing a Flavour
func purchaseFlavourHandler(w http.ResponseWriter, r *http.Request) {
	// Get the flavourID value from the URL parameters
	params := mux.Vars(r)
	flavourID := params["flavourID"]

	// Read the request body
	body, err := io.ReadAll(r.Body)
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

	// Check if the transactions map is nil or not initialized
	if transactions == nil {
		http.Error(w, "Error: no active transactions found.", http.StatusNotFound)
		return
	}

	// Retrieve the transaction from the transactions map
	transaction, exists := transactions[purchase.TransactionID]
	if !exists {
		http.Error(w, "Error: transaction not found", http.StatusNotFound)
		return
	}

	// Check if the transaction is still valid (within 20 seconds for testing)
	elapsedTime := time.Since(transaction.StartTime)
	if elapsedTime > 20*time.Second {
		http.Error(w, "Error: transaction Timeout", http.StatusRequestTimeout)
		delete(transactions, purchase.TransactionID)
		return
	}

	// Perform the Flavour purchase
	// ...

	// Remove the transaction from the transactions map
	delete(transactions, purchase.TransactionID)

	// Delete Flavour that is no longer avaialable (in MongoDB)
	// There will be a function in mongo_manager.go

	// Respond with the FlavourID as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResPur{
		FlavourID: flavourID,
		BuyerID:   purchase.BuyerID,
		Status:    "Completed",
	})
}
