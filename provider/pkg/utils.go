package pkg

import (
	"fmt"
	"time"
)

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
	if selector.CPU != 0 && flavour.Characteristics.CPU < selector.CPU {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.CPUMinimum; i < flavour.Characteristics.CPU; i = i + flavour.Policy.Partitionable.CPUStep {
				if i >= selector.CPU {
					check++
				}
			}
			if check == 0 {
				return false
			}
		}

	}

	// Check if the RAM matches
	if selector.RAM != 0 && flavour.Characteristics.RAM < selector.RAM {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.RAMMinimum; i < flavour.Characteristics.RAM; i = i + flavour.Policy.Partitionable.RAMStep {
				if i >= selector.RAM {
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
	if selector.CPU != 0 && flavour.Characteristics.CPU > selector.CPU {
		if flavour.Policy.Partitionable == nil {
			return false
		} else {
			check := 0
			for i := flavour.Policy.Partitionable.CPUMinimum; i < flavour.Characteristics.CPU; i = i + flavour.Policy.Partitionable.CPUStep {
				if i <= selector.CPU {
					check++
				}
			}
			if check == 0 {
				return false
			}
		}

	}

	// Check if the RAM matches
	if selector.RAM != 0 && flavour.Characteristics.RAM > selector.RAM {
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
