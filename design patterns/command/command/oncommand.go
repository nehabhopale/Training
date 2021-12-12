package command

import   "devices/device"

type OnCommand struct{
	device  device.Device
}

func (o OnCommand)Execute(){
	o.device.On()
}
func NewOnCmd(device device.Device) OnCommand{
	return OnCommand{
		device:device,
	}
}