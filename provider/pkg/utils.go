package pkg

import (
	"fmt"
	"time"
)

// Generates a unique transaction ID using the current timestamp
func generateTransactionID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
