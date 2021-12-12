package wallet

type Wallet struct {
	balance int
}

func (w *Wallet) GetBalance() int {
	return w.balance
}

func (w *Wallet) Credit(amount int) int {
	w.balance = w.balance + amount
	return w.balance
}
func (w *Wallet) Debit(amount int) int {
	w.balance = w.balance - amount
	return w.balance
}
func New(balance int) *Wallet {
	return &Wallet{
		balance: balance,
	}
}
