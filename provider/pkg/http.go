package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartHttpServer starts a new HTTP server
func StartHttpServer() {
	// mux creation
	router := mux.NewRouter()

	// routes definition
	router.HandleFunc("/api/listflavours", listAllFlavoursHandler).Methods("GET")
	router.HandleFunc("/api/listflavours/selector", listAllFlavoursSelectorHandler).Methods("GET")
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
