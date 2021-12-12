package player

import("tictac/mark")

type Player struct{
	name string
	mark mark.Mark
}

func (p Player) GetName()string{
	return p.name
}
func (p Player) GetMark() mark.Mark{
	return p.mark
}

func NewPlayer(playerName string,mark mark.Mark) *Player{
	return &Player{
		name:playerName,
		mark:mark,
	}
	
}