package store

type checkout struct {
	products map[string]product
	cart     []string
}

type product struct {
	name  string
	price euroCent
}

func NewCheckout() *checkout {
	c := new(checkout)
	c.products = map[string]product{
		"VOUCHER": product{name: "Cabify Voucher", price: 500},
		"TSHIRT":  product{name: "Cabify T-Shirt", price: 2000},
		"MUG":     product{name: "Cabify Coffe Mug", price: 750},
	}
	c.cart = []string{}
	return c
}

func (c *checkout) Scan(item string) {
	c.cart = append(c.cart, item)
}

func (c *checkout) GetTotal() Euro {
	total := euroCent(0)
	for _, it := range c.cart {
		total += c.products[it].price
	}

	return total.Euro()
}
