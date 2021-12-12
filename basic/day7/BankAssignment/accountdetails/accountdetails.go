package accountdetails

type Account struct {
	Accholderid uint64
	AccNo       uint64
	Balance     uint64
	Acctype     string
	Accifsc     uint64
}

func (a *Account) GetAccNo() uint64 {
	return a.AccNo
}
func (a *Account) GetBalance() uint64 {
	return a.Balance
}
func (a *Account) GetAccType() string {
	return a.Acctype
}
func (a *Account) SetAccType(newType string) {
	a.Acctype = newType
}
func (a *Account) GetAccifc() uint64 {
	return a.Accifsc
}
func (a *Account) SetAccNo(newAccNo uint64) {
	a.AccNo = newAccNo
}
func (a *Account) SetBalance(newBalance uint64) {
	a.Balance = newBalance
}
func New(accountNo uint64, accountHolderId uint64, balance uint64, accountype string, ifsc uint64) *Account {
	return &Account{
		Accholderid: accountHolderId,
		AccNo:       accountNo,
		Acctype:     accountype,
		Balance:     balance,
		Accifsc:     ifsc,
	}
}
