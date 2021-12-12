package main
import "fmt"
type Car string

func (c Car) Accelerator(){
	fmt.Println("speed up")
}
func (c Car) Brake(){
	fmt.Println("stopping")
}
func (c Car) Steer(direction string){
	fmt.Println(direction)
}
type Truck string

func (T Truck) Accelerator(){
	fmt.Println("speed up")
}
func (T Truck) Brake(){
	fmt.Println("stopping")
}
func (T Truck) Steer(direction string){
	fmt.Println(direction)
}

func (T Truck) LoadCarGo(cargo string){
	fmt.Println(cargo)
}
type Vehicle interface{
	Accelerator()
	Brake()
	Steer(string)
}
func TryVehicle(vehicle Vehicle){
	vehicle.Accelerator()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck,ok:=vehicle.(Truck)
	if ok{
		truck.LoadCarGo("test cargo")
	}
}
func main(){
	var vehicle Vehicle=Car("TOYODA")
	vehicle.Accelerator()
	vehicle.Steer("left")
	vehicle=Truck("fnord")
	vehicle.Brake()
	vehicle.Steer("right")
	TryVehicle(vehicle)
}