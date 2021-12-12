package main

import (
	"acc/walletFacade"
	"fmt"
)

func main() {
	fmt.Println()
	walletFacade := walletFacade.NewWalletFacade(23456789, 1234, 1000)
	fmt.Println()
	walletFacade.AddMoneyToWallet(23456789, 1234, 10)
	walletFacade.DeductMoneyFromWallet(23456789, 1234, 5)

}
