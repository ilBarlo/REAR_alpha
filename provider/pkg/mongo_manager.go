package pkg

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// init used for connecting with mongoDB
func init() {
	// Params to connect
	const uri = "mongodb+srv://ilBarlo:FlavourGenerator@clusterbarlo.qnlqwmd.mongodb.net/?retryWrites=true&w=majority"
	const dbName = "flavours"
	const colName = "rear"

	// Client option
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to mongodb
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection istance is ready")
}

// listAllFlavours gets all the Flavours
func listAllFlavours() ([]Flavour, error) {

	// Define the filter for the query
	filter := bson.M{}

	// Perform the query
	var flavours []Flavour
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	// Iterate through the query results
	for cur.Next(context.Background()) {
		var f Flavour
		if err := cur.Decode(&f); err != nil {
			return nil, err
		}
		flavours = append(flavours, f)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return flavours, nil
}

// findMatchingFlavours find flavours that fit the Selector
func findMatchingFlavours(selector Selector) []Flavour {
	// Placeholder for the list of matching Flavours
	matchingFlavours := []Flavour{}

	availableFlavours, err := listAllFlavours()
	fmt.Println(err)

	// Iterate over the list of available Flavours
	for _, flavour := range availableFlavours {
		// Check if the Flavour matches the selector criteria
		if flavourMatchesSelector(flavour, selector) {
			// Add the matching Flavour to the list
			matchingFlavours = append(matchingFlavours, flavour)
		}
	}

	return matchingFlavours
}

// findMatchingFlavours find flavours that fit the Selector
func findMatchingFlavoursLess(selector Selector) []Flavour {
	// Placeholder for the list of matching Flavours
	matchingFlavours := []Flavour{}

	availableFlavours, err := listAllFlavours()
	fmt.Println(err)

	// Iterate over the list of available Flavours
	for _, flavour := range availableFlavours {
		// Check if the Flavour matches the selector criteria
		if flavourMatchesSelectorLess(flavour, selector) {
			// Add the matching Flavour to the list
			matchingFlavours = append(matchingFlavours, flavour)
		}
	}

	return matchingFlavours
}

// findMatchingFlavours find flavours that fit the Selector
func findMatchingFlavoursMore(selector Selector) []Flavour {
	// Placeholder for the list of matching Flavours
	matchingFlavours := []Flavour{}

	availableFlavours, err := listAllFlavours()
	fmt.Println(err)

	// Iterate over the list of available Flavours
	for _, flavour := range availableFlavours {
		// Check if the Flavour matches the selector criteria
		if flavourMatchesSelectorMore(flavour, selector) {
			// Add the matching Flavour to the list
			matchingFlavours = append(matchingFlavours, flavour)
		}
	}

	return matchingFlavours
}

// getFlavourByID gets a Flavour by ID
func getFlavourByID(flavourID string) (*Flavour, error) {
	// Create a filter for the query
	filter := bson.M{"flavourID": flavourID}

	// Define a variable to store the retrieved Flavour
	var flavour Flavour

	// Perform the MongoDB query
	err := collection.FindOne(context.TODO(), filter).Decode(&flavour)
	if err != nil {
		return nil, err
	}

	return &flavour, nil
}

// getFlavourTypes gets all the types of available Flavour
func getFlavourTypes() []string {
	flavours, _ := listAllFlavours()
	types := make([]string, 0)
	typeMap := make(map[string]bool)

	for _, flavour := range flavours {
		typeMap[flavour.Type] = true
	}

	for t := range typeMap {
		types = append(types, t)
	}

	return types
}
