package pizzamenu

type OnionTopping struct{
	Pizza Pizza
}

func (t *OnionTopping) GetPrize() uint32 {
	pizzaPrice := t.Pizza.GetPrize()
	return pizzaPrice + 20
}