package pizzamenu

type CheeseTopping struct{
	Pizza Pizza
}

func (c *CheeseTopping) GetPrize() uint32 {
	pizzaPrice := c.Pizza.GetPrize()
	return pizzaPrice + 10
}