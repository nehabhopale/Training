package main

 import ("Guitar/guitar"
)

func main(){
	gs1:=guitar.NewGuitarSpecs("model1",guitar.Fender,guitar.Electric,guitar.Sitka,guitar.Sitka,11)
	g1:=guitar.NewGuitar("1",300,gs1)

	gs2:=guitar.NewGuitarSpecs("model2",guitar.Martin,guitar.Electric,guitar.Sitka,guitar.Sitka,11)
	g2:=guitar.NewGuitar("2",300,gs2)
	shop:=guitar.NewInventory() 
	shop.AddGuitar(g1)
	shop.AddGuitar(g2)
	shop.PrintGuitar()
	//new:= shop.GetGuitar("1")
	//fmt.Println(*new)
	///////////////
	k1:=&guitar.GuitarSpecs{}
	k1.SetModel("model1")
	k1.SetBuilder(guitar.Fender)
	k2:=guitar.NewGuitar("1",300,k1)
	shop.SearchGuitar(k2)
	//fmt.Println(ppp)
	

	

	

}