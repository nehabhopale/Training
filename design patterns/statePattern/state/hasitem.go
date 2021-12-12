package state

import (
	"fmt"
)

type HasItemState struct {
	vendingMachine *VendingMachine
}

//we r having item and we r requesting it
func (i *HasItemState) RequestItem() error {
	// if i.vendingMachine.itemCount == 0 {
	// 	i.vendingMachine.SetState(i.vendingMachine.noItem)
	// 	return fmt.Errorf("No item present")
	// }
	fmt.Println("Item requestd by customer")
	i.vendingMachine.SetState(i.vendingMachine.itemRequested)
	return nil
}
func (i *HasItemState) AddItem(count uint8) error {
	if count == 0 {
		return fmt.Errorf("count shouldn't be zero")
	}
	fmt.Println(count, "no of items added")
	i.vendingMachine.IncrementItemCount(count)
	return nil
}
func (i *HasItemState) InsertMoney(money uint16) error {
	return fmt.Errorf("u will need to select item")
}
func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("u will need to select item to be dispended")
}
