package guitar
import "fmt"
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
	Electric WoodType="electric"
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

func (g Guitar)DisplayGuitar(){
	fmt.Println("*********************")
	fmt.Println(g.serialNo)
	fmt.Println(g.price)
	fmt.Println(g.GetSpecs().GetModel())
	fmt.Println(g.GetSpecs().GetBuilder())
	fmt.Println(g.GetSpecs().GetWoodType())
	fmt.Println(g.GetSpecs().GetBackWood())
	fmt.Println(g.GetSpecs().GetFrontWood())
	fmt.Println(g.GetSpecs().GetNoOfStrings())
}