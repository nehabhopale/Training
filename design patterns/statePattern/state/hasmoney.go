package state

import (
	"fmt"
)

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) AddItem(count uint8) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money uint16) error {
	return fmt.Errorf("Item out of stock")
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.hasItem)
	}
	return nil
}
