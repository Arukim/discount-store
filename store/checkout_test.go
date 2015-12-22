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
