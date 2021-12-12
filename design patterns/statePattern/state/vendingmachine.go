package state

import (
	"fmt"
)

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount uint8
	itemPrice uint16
}

func New(itemCount uint8, itemPrice uint16) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItem{
		vendingMachine: v,
	}

	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count uint8) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money uint16) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementItemCount(count uint8) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
