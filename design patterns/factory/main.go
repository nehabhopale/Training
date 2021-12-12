package main
import "factory/factory"

//provides a way to hide the creation logic of the instances being created.
//The client only interacts with a factory struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding concrete structs and returns the correct instance back.

func main(){
	newCar:=factory.Make(factory.Bmw)  //interacting with factory struct and tells bmw instance needed
	newCar.Start()
	newCar.Stop()
	newCar1:=factory.Make(factory.Tesla)  //interacting with factory struct and tells bmw instance needed
	newCar1.Start()
	newCar1.Stop()
}