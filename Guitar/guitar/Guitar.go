package guitar

type Builder string
const(
	Fender Builder ="fender"
	Martin Builder ="martin"
	Gibson Builder ="gibson"
)

type Wood string

const(
	Mahagony Wood ="mahagony"
	Maple Wood ="maple"
	Cocobolo Wood="cocobolo"
	Cedar Wood ="cedar"
	Sitka Wood="sitka"
)

type WoodType string

const(
	Accoustic WoodType="accoustic"
	Electric WoodType="accoustic"
)

type Guitar struct{
	serialNo string
	price uint16
	specs *GuitarSpecs
}
func (g *Guitar)GetSerialNo()string{
	return g.serialNo
}
func (g *Guitar)GetPrice()uint16{
	return g.price
}
func (g *Guitar)GetSpecs()*GuitarSpecs{
	return g.specs
}

func NewGuitar(serialNo string,price uint16,specs *GuitarSpecs)*Guitar{
	return &Guitar{
		serialNo :serialNo,
		price :price,
		specs :specs,
	}
}

