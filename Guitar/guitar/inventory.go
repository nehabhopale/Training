package guitar
import ("fmt"
)

type Inventory struct{
	allGuitars []*Guitar
}

func (I *Inventory)AddGuitar(g *Guitar){
	I.allGuitars=append(I.allGuitars,g)
}
func(I *Inventory) RemoveGuitar(no string ){
	I.allGuitars=DelFromGuitar(I.allGuitars,no)
}
func DelFromGuitar(guitarList []*Guitar,no string)[]*Guitar{
	subLength:=len(guitarList)
	for pos ,guitar:=range guitarList{
		if guitar.GetSerialNo()==no{
			 lastIdx:=subLength-1
			 guitarList[pos]=guitarList[lastIdx]
			 guitarList=guitarList[:pos]
			 return guitarList
		}
	}
	return nil
}
func (I *Inventory)GetGuitar(no string)*Guitar{
	for _,guitar:=range(I.allGuitars){
		if guitar.GetSerialNo()==no{
			return guitar
		}
	}
	return nil
}

func contains(s []interface{}, val interface{}) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}

	return false
}

func(I *Inventory)SearchGuitar( g *Guitar){
	searchSpecs:=g.GetSpecs()
	specList:=MatchGuitar(searchSpecs)
	fmt.Println(specList)
	 var resultGuitar []Guitar
	
	for _,guitar:=range(I.allGuitars){
		fmt.Println(guitar)
		i:=0
		j:=0

		for i<len(specList){
			
			if specList[i]==guitar.GetSpecs().GetModel()||specList[i]==guitar.GetSpecs().GetBuilder()||specList[i]==guitar.GetSpecs().GetWoodType()||specList[i]==guitar.GetSpecs().GetBackWood()||specList[i]==guitar.GetSpecs().GetFrontWood()||specList[i]==guitar.GetSpecs().GetNoOfStrings(){
				j++

			}
			i++

		}
		if j==len(specList){
			resultGuitar=append(resultGuitar,*guitar)

		}
		fmt.Println(j)
			
	}
	for _,guitar:=range(resultGuitar){
		guitar.DisplayGuitar()
	}

}






func(I *Inventory)PrintGuitar(){
	for _,guitar:=range(I.allGuitars){
		fmt.Println(*guitar)
	}
}
func NewInventory() *Inventory {
	return &Inventory{}
}