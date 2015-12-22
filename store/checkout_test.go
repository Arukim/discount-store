package store

import "testing"

type checkoutPair struct {
	cart  []string
	total Euro
}

var tests = []checkoutPair{
	{[]string{"VOUCHER", "TSHIRT", "MUG"}, Euro(32.50)},
	{[]string{"VOUCHER", "TSHIRT", "VOUCHER"}, Euro(25.00)},
	{[]string{"TSHIRT", "TSHIRT", "TSHIRT", "VOUCHER", "TSHIRT"}, Euro(81.00)},
	{[]string{"VOUCHER", "TSHIRT", "VOUCHER", "VOUCHER", "MUG", "TSHIRT", "TSHIRT"}, Euro(74.50)},
}

func TestAllBaseCheckouts(t *testing.T) {
	for i, testCase := range tests {
		checkout := NewCheckout()

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

func TestSecondCheckout(t *testing.T) {
}

func TestEmptyCart(t *testing.T) {
	totalExpected := Euro(0)
	checkout := NewCheckout()
	total := checkout.GetTotal()
	if total != totalExpected {
		t.Errorf("Expected %v, go %v", totalExpected, total)
	}
}

func TestUnknownProduct(t *testing.T) {
	checkout := NewCheckout()
	err := checkout.Scan("IPhone9.5")
	if err == nil {
		t.Error("Currently no mobile phones are sold")
	}
}
