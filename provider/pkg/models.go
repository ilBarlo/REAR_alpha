package pkg

import "time"

// Flavour represents a Flavour object with its characteristics and policies.
type Flavour struct {
	FlavourID       string          `json:"flavourID"`
	ProviderID      string          `json:"providerID"`
	Type            string          `json:"type"`
	Characteristics Characteristics `json:"characteristics"`
	Policy          Policy          `json:"policy"`
	Owner           Owner           `json:"owner"`
	Price           Price           `json:"price"`
	ExpirationTime  time.Time       `json:"expirationTime"`
	OptionalFields  OptionalFields  `json:"optionalFields"`
}

// Characteristics represents the characteristics of a Flavour, such as CPU and RAM.
type Characteristics struct {
	CPU int `json:"cpu"`
	RAM int `json:"ram"`
}

// Policy represents the policy associated with a Flavour, which can be either Partitionable or Aggregatable.
type Policy struct {
	Partitionable *Partitionable `json:"partitionable,omitempty"`
	Aggregatable  *Aggregatable  `json:"aggregatable,omitempty"`
}

// Partitionable represents the partitioning properties of a Flavour, such as the minimum and incremental values of CPU and RAM.
type Partitionable struct {
	CPUMinimum int `json:"cpuMinimum"`
	RAMMinimum int `json:"ramMinimum"`
	CPUStep    int `json:"cpuStep"`
	RAMStep    int `json:"ramStep"`
}

// Aggregatable represents the aggregation properties of a Flavour, such as the minimum instance count.
type Aggregatable struct {
	MinCount int `json:"minCount"`
	MaxCount int `json:"maxCount"`
}

// Owner represents the owner of a Flavour, with associated ID, IP, and domain name.
type Owner struct {
	ID         string `json:"ID"`
	IP         string `json:"IP"`
	DomainName string `json:"domainName"`
}

// Price represents the price of a Flavour, with the amount, currency, and period associated.
type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Period   string  `json:"period"`
}

// OptionalFields represents the optional fields of a Flavour, such as availability.
type OptionalFields struct {
	Availability bool `json:"availability,omitempty"`
}

// Selector represents the criteria for selecting Flavours.
type Selector struct {
	CPU         int    `json:"cpu,omitempty"`
	RAM         int    `json:"ram,omitempty"`
	Policy      Policy `json:"policy,omitempty"`
	FlavourType string `json:"type,omitempty"`
}

// getSyntat returns the syntax for the selector
func getSyntax() Selector {

	// Create the selector
	selector := Selector{
		CPU:         2,
		RAM:         4,
		FlavourType: "VM",
	}

	return selector
}
