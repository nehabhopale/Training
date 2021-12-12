package state

import (
	"fmt"
)

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) AddItem(count uint8) error {
	return fmt.Errorf("item dispense will proceed")

}
func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already added")

}
func (i *ItemRequestedState) InsertMoney(money uint16) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.SetState(i.vendingMachine.hasMoney)
	return nil
}
func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("add money")

}
