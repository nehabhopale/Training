package device

import "fmt"

type Mobile struct{
	isRunning bool
}
func NewMobile(isRunning bool) Mobile{
	return Mobile{
		isRunning:isRunning,
	}
}
func (m Mobile) On() {
	if m.isRunning {
		fmt.Println("mobile is already in on state")
	}else{
		fmt.Println("mobile is turned on ")
	}
}
func (m Mobile)Off(){
	if m.isRunning{
		fmt.Println("mobile is turned off")
	}else{
		fmt.Println("mobile is already off")

	}
}
