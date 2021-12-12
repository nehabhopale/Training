package factory
import "fmt"
type tesla struct{
	modelName string
}

func NewTesla(modelName string) tesla{
	return tesla{
		modelName:modelName,
	}
}
func (t tesla) Start(){
	fmt.Println("tesla  started")
}
func (t tesla) Stop(){
	fmt.Println("tesla stopped")
}