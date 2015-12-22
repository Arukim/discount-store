package store

import "fmt"

// For internal calculations
type euroCent int

// Only input/output
type Euro float64

func (euro Euro) String() string {
	return fmt.Sprintf("%.2f€", euro)
}
