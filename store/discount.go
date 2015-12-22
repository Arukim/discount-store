package store

type Discounter interface {
	Discount(c *checkout) chan bool
}

type discounterA struct {
	itemCode string
}

func NewDiscounterA(code string) *discounterA {
	d := discounterA{itemCode: code}
	return &d
}

func (d discounterA) Discount(c *checkout) chan bool {
	b := make(chan bool)
	go func() {
		counter := 0
		for i := range c.cart {
			it := &c.cart[i]
			if it.code == d.itemCode {
				counter++
				if counter%2 == 0 {
					it.price = 0
				}
			}
		}
		b <- true
	}()
	return b
}

type discounterB struct {
	itemCode string
}

func NewDiscounterB(code string) *discounterB {
	d := discounterB{itemCode: code}
	return &d
}

func (d discounterB) Discount(c *checkout) chan bool {
	b := make(chan bool)
	go func() {
		counter := 0
		for _, it := range c.cart {
			if it.code == d.itemCode {
				counter++
			}
		}

		if counter >= 3 {
			for i, it := range c.cart {
				if it.code == d.itemCode {
					c.cart[i].price = euroCent(1900)
				}
			}
		}

		b <- true
	}()
	return b
}
