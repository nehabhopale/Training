package main

import(
	"devices/device"
	"devices/button"
	"devices/command"
)


// Command pattern is used when we want to execute commands(ON, OFF in this case) based on different objects (TV, mobile in this case)
func main(){
	Samsung:=device.NewMobile(true)
	yellowButton:=button.NewButton(command.NewOnCmd(Samsung))
	yellowButton.Press()

	DishTv:=device.NewTv(false)
	greenButton:=button.NewButton(command.NewOffCmd(DishTv))
	greenButton.Press()
}