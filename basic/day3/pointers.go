package main

import "fmt"

func main() {

	var a int = 10
	fmt.Println("value of a", a)    //10
	fmt.Println("address of a", &a) //addrees of a

	var p *int = &a                                     //p of type pointer to datatype of a * used for deref
	fmt.Println("value of p", p)                        //address of a is value of a
	fmt.Println("address of a", &p)                     // address of p
	fmt.Println("value at the address stored by p", *p) // derefrencing of p

	var p2p *(*int) = &p                                   //var of type pointer point to pointer of datatype type of p
	fmt.Println("value of p2p", p2p)                       //address of a is value of p
	fmt.Println("address of p2p", &p2p)                    // address of p2p
	fmt.Println("value at the address stored by p", *p2p)  // derefrencing of p := value of p2p:=add of a
	fmt.Println("value at the address stored by p", **p2p) //value at the address &p2p

}
