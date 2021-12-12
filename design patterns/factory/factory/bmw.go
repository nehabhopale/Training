package factory
import "fmt"
type bmw struct{
	modelName string
}

func NewBmw(modelName string) bmw{
	return bmw{
		modelName:modelName,
	}
}
func (b bmw) Start(){
	fmt.Println("bmw started")
}
func (b bmw) Stop(){
	fmt.Println("bmw stopped")
}