package state

import (
	"fmt"
)

type NoItem struct {
	vendingMachine *VendingMachine
}

func (i *NoItem) RequestItem() error {
	return fmt.Errorf("Item is not available")
}

func (i *NoItem) AddItem(count uint8) error {
	i.vendingMachine.IncrementItemCount(count)
	i.vendingMachine.SetState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItem) InsertMoney(money uint16) error {
	return fmt.Errorf("Sorry Item is not available")
}
func (i *NoItem) DispenseItem() error {
	return fmt.Errorf("sorry item is not available so cant be dispensed")
}
