package button

import "devices/command"

type Button struct{
	 command command.Command
}

func NewButton(command command.Command) Button{
	return Button{
		command:command,
	}
}

func (b Button)Press(){
	b.command.Execute()
}