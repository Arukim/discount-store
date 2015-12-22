package store

type checkout struct {
}

func NewCheckout() *checkout {
	return &checkout{}
}

func (c *checkout) Scan(item string) {

}

func (c *checkout) GetTotal() Euro {
	return Euro(10)
}
