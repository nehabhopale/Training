package factory

import ("factory/automobile")

type car string

const(
	Bmw car ="bmw"
	Mercedes car ="mercedes"
	Tesla car ="tesla"
)

func Make( brand car) automobile.Automobile{
	switch brand{
	case "bmw":
		return NewBmw("Bmw")
	case "mercedes":
		return NewMercedes("mercedes")
	case "tesla":
		return NewTesla("tesla")
	}
	return nil

}