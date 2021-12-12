 package main
 import "fmt"
// import ("reflect"
// "fmt")

// type Builder string
// const(
// 	Fender Builder ="fender"
// 	Martin Builder ="martin"
// 	Gibson Builder ="gibson"
// )


// type GuitarSpecs struct{
// 	//model string
// 	builder Builder
// }

// func (g GuitarSpecs)ap()[]interface{}{
// 	var l []interface{}
// 	fmt.Println(reflect.TypeOf(g.builder))
// 	l=append(l,reflect.TypeOf(g.builder).String())
// 	fmt.Println(l)
// 	return l

// }
// func (g GuitarSpecs)get()string{
// 	return reflect.TypeOf(g.builder).String()
// }

// // func main()
// 	g:= GuitarSpecs{builder:Fender}
// 	l:=g.ap()
// 	fmt.Println(g.get())
// 	for _,v :=range(l){
// 		if v==g.get(){
// 			fmt.Println("true")
// 		}
// 	}

// }

func main(){
	var cells [3][3]string
	
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v", cells[i][j])
		}
		fmt.Println()
	}
}
