package store

import "testing"

func TestFirstCheckout(t *testing.T) {
	totalExpected := Euro(32.50)
	checkout := NewCheckout()

	checkout.Scan("VOUCHER")
	checkout.Scan("TSHIRT")
	checkout.Scan("MUG")

	total := checkout.GetTotal()
	if total != totalExpected {
		t.Errorf("Expected %v, got %v", totalExpected, total)
	}
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
