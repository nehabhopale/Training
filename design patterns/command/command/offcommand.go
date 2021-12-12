package command

import   "devices/device"

type OffCommand struct{
	device  device.Device
}

func (o OffCommand)Execute(){
	o.device.Off()
}

func NewOffCmd(device device.Device) OffCommand{
	return OffCommand{
		device:device,
	}
}