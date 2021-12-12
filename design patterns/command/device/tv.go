package device

import "fmt"
type Tv struct{
	isRunning bool
}
func NewTv(isRunning bool) *Tv{
	return &Tv{
		isRunning:isRunning,
	}
}
func (t *Tv) On() {
	if t.isRunning {
		fmt.Println("Tv is already in on state")
	}else{
		fmt.Println("Tv is turned on ")
	}
}
func (t *Tv)Off(){
	if t.isRunning{
		fmt.Println("Tv is turned off")
	}else{
		fmt.Println("Tv is already off")

	}
}
