package store

import "fmt"

// For internal calculations
type euroCent int

// Only input/output
type Euro float64

// pretty print Euro
func (euro Euro) String() string {
	return fmt.Sprintf("%.2f€", euro)
}

// euroCent to Euro conversion
func (cent euroCent) Euro() Euro {
	return (Euro)(cent) / 100
}

func (euro Euro) EuroCent() euroCent {
	return (euroCent)(euro * 100)
}
