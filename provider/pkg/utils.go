package pkg

import (
	"fmt"
	"time"
)

func init() {
	types = populateTypes()
}

// generateTransactionID Generates a unique transaction ID using the current timestamp
func generateTransactionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// flavourMatchesSelector checks if the selector match some flavours
func flavourMatchesSelector(flavour Flavour, selector Selector) bool {
	// Check if the CPU matches
	if selector.CPU != 0 && flavour.Characteristics.CPU != selector.CPU {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.CPUMinimum; i < flavour.Characteristics.CPU; i = i + flavour.Policy.Partitionable.CPUStep {
				if i == selector.CPU {
					check++
				}
			}
			if check == 0 {
				return false
			}
		}

	}

	// Check if the RAM matches
	if selector.RAM != 0 && flavour.Characteristics.RAM != selector.RAM {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.RAMMinimum; i < flavour.Characteristics.RAM; i = i + flavour.Policy.Partitionable.RAMStep {
				if i == selector.RAM {
					check++
				}
			}
			if check == 0 {
				return false
			}
		}
	}

	// Check if the type matches
	if selector.FlavourType != "" && flavour.Type != selector.FlavourType {
		return false
	}

	return true
}

// flavourMatchesSelector checks if the selector match some flavours
func flavourMatchesSelectorMore(flavour Flavour, selector Selector) bool {
	// Check if the CPU matches
	if selector.MoreThanCPU != 0 && flavour.Characteristics.CPU < selector.MoreThanCPU {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.CPUMinimum; i < flavour.Characteristics.CPU; i = i + flavour.Policy.Partitionable.CPUStep {
				if i >= selector.MoreThanCPU {
					check++
				}
			}
			if check == 0 {
				return false
			}
		}

	}

	// Check if the RAM matches
	if selector.MoreThanRAM != 0 && flavour.Characteristics.RAM < selector.MoreThanRAM {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.RAMMinimum; i < flavour.Characteristics.RAM; i = i + flavour.Policy.Partitionable.RAMStep {
				if i >= selector.MoreThanRAM {
					check++
				}
			}
			if check == 0 {
				return false
			}
		}
	}

	// Check if the type matches
	if selector.FlavourType != "" && flavour.Type != selector.FlavourType {
		return false
	}

	return true
}

// flavourMatchesSelector checks if the selector match some flavours
func flavourMatchesSelectorLess(flavour Flavour, selector Selector) bool {
	// Check if the CPU matches
	if selector.LessThanCPU != 0 && flavour.Characteristics.CPU > selector.LessThanCPU {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.CPUMinimum; i < flavour.Characteristics.CPU; i = i + flavour.Policy.Partitionable.CPUStep {
				if i <= selector.LessThanCPU {
					check++
				}
			}
			if check == 0 {
				return false
			}
		}

	}

	// Check if the RAM matches
	if selector.LessThanRAM != 0 && flavour.Characteristics.RAM > selector.LessThanRAM {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.RAMMinimum; i < flavour.Characteristics.RAM; i = i + flavour.Policy.Partitionable.RAMStep {
				if i <= selector.RAM {
					check++
				}
			}
			if check == 0 {
				return false
			}
		}
	}

	// Check if the type matches
	if selector.FlavourType != "" && flavour.Type != selector.FlavourType {
		return false
	}

	return true
}

// populateTypes creates an array of Type objects and initializes each object with the FlavourType and Description values (testing only)
func populateTypes() []Type {
	types := []Type{
		{
			FlavourType: "k8s",
			Description: "Flavour type for Kubernetes-based deployments",
		},
		{
			FlavourType: "VM",
			Description: "Flavour type for virtual machine-based deployments",
		},
		{
			FlavourType: "AWS VM",
			Description: "Flavour type for Amazon Web Services (AWS) virtual machine-based deployments",
		},
	}
	return types
}
