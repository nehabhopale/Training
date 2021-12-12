package factory
import "fmt"
type mercedes struct{
	modelName string
}

func NewMercedes(modelName string) mercedes{
	return mercedes{
		modelName:modelName,
	}
}
func (m mercedes) Start(){
	fmt.Println("mercedes started")
}
func (m mercedes) Stop(){
	fmt.Println("mercedes stopped")
}