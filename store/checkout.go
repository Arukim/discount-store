package store

import (
	"errors"
)

type checkout struct {
	catalog     map[string]product
	cart        []product
	discounters []Discounter
}

type product struct {
	code  string
	name  string
	price euroCent
}

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

func (c *checkout) Scan(item string) error {
	if it, ok := c.catalog[item]; ok {
		c.cart = append(c.cart, product{code: it.code, price: it.price})
		return nil
	}
	return errors.New("no such item in store")
}

func (c *checkout) GetTotal() Euro {
	total := euroCent(0)

	for _, disc := range c.discounters {
		b := disc.Discount(c)
		<-b
	}

	for _, it := range c.cart {
		total += it.price
	}

	return total.Euro()
}
