package walletFacade

import (
	"acc/account"
	"acc/ledger"
	"acc/notification"
	"acc/securityCode"
	"acc/wallet"
	"fmt"
)

type walletFacade struct {
	account      *account.Account
	wallet       *wallet.Wallet
	securityCode *securityCode.SecurityCode
	notification *notification.Notification
	ledger       *ledger.Ledger
}

func NewWalletFacade(accountNo uint64, code uint16, balance int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      account.New(accountNo),
		securityCode: securityCode.New(code),
		wallet:       wallet.New(balance),
		notification: &notification.Notification{},
		ledger:       &ledger.Ledger{},
	}
	fmt.Println("Account created")
	return walletFacacde
}
func (w *walletFacade) AddMoneyToWallet(accNo uint64, securityCode uint16, amount int) {
	fmt.Println("Starting add money to wallet")

	if (w.account.CheckAccount(accNo)) && (w.securityCode.CheckCode(securityCode)) {
		w.wallet.Credit(amount)
		w.notification.SendWalletCreditNotification()
		w.ledger.MakeEntry(accNo, "credit", amount)
		
	}
	fmt.Println(w.wallet.GetBalance())

}
func (w *walletFacade) DeductMoneyFromWallet(accNo uint64, securityCode uint16, amount int) {
	fmt.Println("Starting debit money from wallet")

	if (w.account.CheckAccount(accNo)) && (w.securityCode.CheckCode(securityCode)) {
		w.wallet.Debit(amount)

		w.notification.SendWalletDebitNotification()
		w.ledger.MakeEntry(accNo, "credit", amount)
	}
	fmt.Println(w.wallet.GetBalance())

}
