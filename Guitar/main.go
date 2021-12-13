package main

 import ("Guitar/guitar"
)

func main(){
	gs1:=guitar.NewGuitarSpecs("model1",guitar.Fender,guitar.Electric,guitar.Sitka,guitar.Sitka,11)
	g1:=guitar.NewGuitar("1",300,gs1)

	gs2:=guitar.NewGuitarSpecs("model2",guitar.Martin,guitar.Electric,guitar.Sitka,guitar.Sitka,1)
	g2:=guitar.NewGuitar("2",300,gs2)

	gs3:=guitar.NewGuitarSpecs("model3",guitar.Martin,guitar.Accoustic,guitar.Mahagony,guitar.Sitka,12)
	g3:=guitar.NewGuitar("3",100,gs3)

	gs4:=guitar.NewGuitarSpecs("model4",guitar.Gibson ,guitar.Electric,guitar.Maple,guitar.Sitka,10)
	g4:=guitar.NewGuitar("4",800,gs4)

	gs5:=guitar.NewGuitarSpecs("model5",guitar.Fender,guitar.Accoustic,guitar.Cocobolo ,guitar.Maple,15)
	g5:=guitar.NewGuitar("5",900,gs5)

	gs6:=guitar.NewGuitarSpecs("model6",guitar.Gibson ,guitar.Accoustic,guitar.Cedar,guitar.Sitka,17)
	g6:=guitar.NewGuitar("6",3000,gs6)

	gs7:=guitar.NewGuitarSpecs("model7",guitar.Fender,guitar.Electric,guitar.Sitka,guitar.Cocobolo ,8)
	g7:=guitar.NewGuitar("7",300,gs7)

	gs8:=guitar.NewGuitarSpecs("model8",guitar.Martin,guitar.Electric,guitar.Mahagony,guitar.Sitka,9)
	g8:=guitar.NewGuitar("8",900,gs8)

	gs9:=guitar.NewGuitarSpecs("model9",guitar.Gibson,guitar.Accoustic,guitar.Maple,guitar.Cedar,1)
	g9:=guitar.NewGuitar("9",400,gs9)

	gs10:=guitar.NewGuitarSpecs("model10",guitar.Fender,guitar.Accoustic,guitar.Sitka,guitar.Sitka,31)
	g10:=guitar.NewGuitar("10",700,gs10)

	gs11:=guitar.NewGuitarSpecs("model11",guitar.Martin,guitar.Electric,guitar.Cedar,guitar.Mahagony,5)
	g11:=guitar.NewGuitar("11",200,gs11)

	gs12:=guitar.NewGuitarSpecs("model12",guitar.Gibson,guitar.Electric,guitar.Cocobolo,guitar.Maple,6)
	g12:=guitar.NewGuitar("12",300,gs12)

	gs13:=guitar.NewGuitarSpecs("model13",guitar.Fender,guitar.Accoustic,guitar.Cocobolo,guitar.Cocobolo,7)
	g13:=guitar.NewGuitar("13",800,gs13)

	gs14:=guitar.NewGuitarSpecs("model14",guitar.Martin,guitar.Accoustic,guitar.Sitka,guitar.Cocobolo,9)
	g14:=guitar.NewGuitar("14",500,gs14)

	gs15:=guitar.NewGuitarSpecs("model15",guitar.Gibson,guitar.Electric,guitar.Cedar,guitar.Sitka,8)
	g15:=guitar.NewGuitar("2",300,gs15)
	///////////////////////////////////////////////
	shop:=guitar.NewInventory() 
	shop.AddGuitar(g1)
	shop.AddGuitar(g2)
	shop.AddGuitar(g3)
	shop.AddGuitar(g4)
	shop.AddGuitar(g5)
	shop.AddGuitar(g6)
	shop.AddGuitar(g7)
	shop.AddGuitar(g8)
	shop.AddGuitar(g9)
	shop.AddGuitar(g10)
	shop.AddGuitar(g11)
	shop.AddGuitar(g12)
	shop.AddGuitar(g13)
	shop.AddGuitar(g14)
	shop.AddGuitar(g15)

	shop.PrintGuitar()
	//new:= shop.GetGuitar("1")
	//fmt.Println(*new)
	///////////////
	k1:=&guitar.GuitarSpecs{}
	k1.SetModel("model1")
	k1.SetBuilder(guitar.Fender)
	k2:=guitar.NewGuitar("1",300,k1)
	shop.SearchGuitar(k2)
	
	

	

	

}