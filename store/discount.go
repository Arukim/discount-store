package store

// Discounter interface and implementations

type Discounter interface {
	// Discount calculation can be long
	Discount(c *checkout) chan bool
}

// Every Nth item you get for free, e.g.
// N = 2, take 1 - pay 1, take 2 - pay 1, take 3 - pay 2, take 4 - pay 2
// N = 3, take 2 - pay 2, take 3 - pay 2, etc...
type discounterA struct {
	itemCode string
	count    int
}

// get new discounterA instance
func NewDiscounterA(code string, count int) *discounterA {
	d := discounterA{itemCode: code, count: count}
	return &d
}

// impl Discounter
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

// Get new price on all products if buying more than X
type discounterB struct {
	itemCode      string
	minCount      int
	discountPrice euroCent
}

// get new discounterB instance
func NewDiscounterB(code string, minCount int, discountPrice Euro) *discounterB {
	d := discounterB{itemCode: code, minCount: minCount, discountPrice: discountPrice.EuroCent()}
	return &d
}

//impl Discounter
func (d discounterB) Discount(c *checkout) chan bool {
	b := make(chan bool)
	go func() {
		counter := 0
		// count up itemCode items in cart
		for _, it := range c.cart {
			if it.code == d.itemCode {
				counter++
			}
		}
		// if have enough - apply discount
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
