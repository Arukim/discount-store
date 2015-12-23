// Usage - create NewCheckout, add some items with Scan, GetTotal in the end
// You can't add items after checkout is finished
// You can add custom disconters, see discount.go
package store

import (
	"errors"
)

type checkout struct {
	catalog     map[string]product
	cart        []product
	discounters []Discounter
	isClosed    bool
	total       euroCent
}

type product struct {
	code  string
	name  string
	price euroCent
}

// get new Checkout instance
func NewCheckout(d []Discounter) *checkout {
	c := new(checkout)
	c.catalog = map[string]product{
		"VOUCHER": product{code: "VOUCHER", name: "Cabify Voucher", price: 500},
		"TSHIRT":  product{code: "TSHIRT", name: "Cabify T-Shirt", price: 2000},
		"MUG":     product{code: "MUG", name: "Cabify Coffe Mug", price: 750},
	}
	c.cart = []product{}
	c.discounters = d
	return c
}

// Add item to checkout
func (c *checkout) Scan(item string) error {
	// Can't add items to already closed checkout
	if c.isClosed {
		return errors.New("checkout is already closed")
	}
	// if item is on catalog - push it into cart
	if it, ok := c.catalog[item]; ok {
		c.cart = append(c.cart, product{code: it.code, price: it.price})
		return nil
	}
	return errors.New("no such item in store")
}

// count total
func (c *checkout) GetTotal() Euro {
	// if total is already counted - return counted value
	if c.isClosed {
		return c.total.Euro()
	}

	// run all discounters
	for _, disc := range c.discounters {
		b := disc.Discount(c)
		<-b
	}

	// get cart sum
	for _, it := range c.cart {
		c.total += it.price
	}

	// set closed flag
	c.isClosed = true
	return c.total.Euro()
}
