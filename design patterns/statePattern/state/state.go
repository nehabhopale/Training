package state

type State interface {
	AddItem(uint8) error
	RequestItem() error
	InsertMoney(money uint16) error
	DispenseItem() error
}
