package store

type Discounter interface {
	Discount(c *checkout) chan bool
}

type discounterA struct {
	itemCode string
	count    int
}

func NewDiscounterA(code string, count int) *discounterA {
	d := discounterA{itemCode: code, count: count}
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
				if counter%d.count == 0 {
					it.price = 0
				}
			}
		}
		b <- true
	}()
	return b
}

type discounterB struct {
	itemCode      string
	minCount      int
	discountPrice euroCent
}

func NewDiscounterB(code string, minCount int, discountPrice euroCent) *discounterB {
	d := discounterB{itemCode: code, minCount: minCount, discountPrice: discountPrice}
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

		if counter >= d.minCount {
			for i, it := range c.cart {
				if it.code == d.itemCode {
					c.cart[i].price = d.discountPrice
				}
			}
		}

		b <- true
	}()
	return b
}
