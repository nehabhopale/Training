package guitar

type GuitarSpecs struct{
	model string
	builder Builder
	woodType WoodType
	backWood Wood
	frontWood Wood
	noOfStrings uint8
}
func NewGuitarSpecs(model string,builder Builder,woodType WoodType,backWood Wood,frontWood Wood,noOfStrings uint8) *GuitarSpecs{
	return &GuitarSpecs{
		model :model,
		builder :builder,
		woodType :woodType,
		backWood :backWood,
		frontWood :frontWood,
		noOfStrings :noOfStrings,
	}
}

func (g *GuitarSpecs)GetModel()string{
	return g.model
}
func (g *GuitarSpecs)GetBuilder()Builder{
	return g.builder

}
func (g *GuitarSpecs)GetWoodType()WoodType{
	return g.woodType
}
func (g *GuitarSpecs)GetBackWood()Wood{
	return g.backWood
}
func (g *GuitarSpecs)GetFrontWood()Wood{
	return g.frontWood
}
func (g *GuitarSpecs)GetNoOfStrings()uint8{
	return g.noOfStrings
}

func (g *GuitarSpecs)SetModel(model string){
	g.model=model
}
func (g *GuitarSpecs)SetBuilder(builder Builder){
	g.builder=builder
}
func (g *GuitarSpecs)SetWoodType(woodType WoodType){
	 g.woodType=woodType
}
func (g *GuitarSpecs)SetBackWood(backwood Wood){
	g.backWood=backwood
}
func (g *GuitarSpecs)SetFrontWood( frontwood Wood){
	 g.frontWood=frontwood
}
func (g *GuitarSpecs)SetNoOfStrings( no uint8){
	g.noOfStrings=no
}

func MatchGuitar(searchGuitar *GuitarSpecs)[]interface{}{
	var param []interface{}
	if searchGuitar.model!=""{
		param=append(param,searchGuitar.GetModel())
	}
	if searchGuitar.builder!=""{
		param=append(param,searchGuitar.GetBuilder())
	}
	if searchGuitar.woodType!=""{
		param=append(param,searchGuitar.GetWoodType())
	}
	if searchGuitar.backWood !=""{
		param=append(param,searchGuitar.GetBackWood())
	}
	if searchGuitar.frontWood !=""{
		param=append(param,searchGuitar.GetFrontWood())
	}
	if searchGuitar.noOfStrings !=0{
		param=append(param,searchGuitar.GetNoOfStrings())
	}

	return (param)

}	