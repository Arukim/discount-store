package store

import "testing"

type checkoutPair struct {
	cart  []string
	total Euro
}

var pricingRules = []Discounter{
	NewDiscounterA("VOUCHER", 2),
	NewDiscounterB("TSHIRT", 3, Euro(19.00)),
}

var tests = []checkoutPair{
	{[]string{"VOUCHER", "TSHIRT", "MUG"}, Euro(32.50)},
	{[]string{"VOUCHER", "TSHIRT", "VOUCHER"}, Euro(25.00)},
	{[]string{"TSHIRT", "TSHIRT", "TSHIRT", "VOUCHER", "TSHIRT"}, Euro(81.00)},
	{[]string{"VOUCHER", "TSHIRT", "VOUCHER", "VOUCHER", "MUG", "TSHIRT", "TSHIRT"}, Euro(74.50)},
}

// Run all base testCases
func TestAllBaseCheckouts(t *testing.T) {
	for i, testCase := range tests {
		checkout := NewCheckout(pricingRules)

		for _, it := range testCase.cart {
			checkout.Scan(it)
		}

		total := checkout.GetTotal()
		if total != testCase.total {
			t.Errorf("In testCase %v Expected %v, got %v",
				i, testCase.total, total)
		}
	}
}

// Empty cart should return zero result
func TestEmptyCart(t *testing.T) {
	totalExpected := Euro(0)
	checkout := NewCheckout(pricingRules)
	total := checkout.GetTotal()
	if total != totalExpected {
		t.Errorf("Expected %v, go %v", totalExpected, total)
	}
}

// UnknownProducts can't be scanned
func TestUnknownProduct(t *testing.T) {
	checkout := NewCheckout(pricingRules)
	err := checkout.Scan("IPhone9.5")
	if err == nil {
		t.Error("Currently no mobile phones are sold")
	}
}

// First GetTotal should close checkout, you can't Scan anything after that
func TestCanPerformOnlyOneCheckout(t *testing.T) {
	checkout := NewCheckout(pricingRules)
	checkout.Scan("VOUCHER")
	checkout.GetTotal()
	err := checkout.Scan("VOUCHER")
	if err == nil {
		t.Error("Can't add item after calling getTotal")
	}
}

// First GetTotal should close checkout and save total price
func TestMultipleGetTotal(t *testing.T) {
	expected := Euro(25)
	checkout := NewCheckout(pricingRules)
	checkout.Scan("VOUCHER")
	checkout.Scan("TSHIRT")
	checkout.Scan("VOUCHER")
	totalA := checkout.GetTotal()
	totalB := checkout.GetTotal()
	if totalA != expected || totalB != expected {
		t.Error("Both totals should be the same")
	}

}
