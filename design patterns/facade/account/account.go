package account

type Account struct {
	accNo uint64
}

func (a Account) CheckAccount(accNo uint64) bool {
	return a.accNo == accNo

}
func New(accNo uint64) *Account {
	return &Account{
		accNo: accNo,
	}
}
