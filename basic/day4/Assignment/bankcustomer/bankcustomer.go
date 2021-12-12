package bankcustomer

import (
	Err "bankinfo/MyError"
	acc "bankinfo/accountdetails"
	"fmt"
)

type customer struct {
	id         uint64
	name       string
	accDetails []*acc.Account
}

func New(id uint64, name string, accountNo uint64, balance uint64, accType string, accountifsc uint64) *customer {
	return &customer{
		id:         id,
		name:       name,
		accDetails: []*acc.Account{acc.New(accountNo, id, balance, accType, accountifsc)},
	}
}

func (c *customer) GetId() uint64 {
	return c.id
}
func (c *customer) GetName() string {
	return c.name
}
func (c *customer) SetId(newId uint64) {
	c.id = newId
}
func (c *customer) SetName(newName string) {
	c.name = newName
}
func (c *customer) GetCustomerDetails() {
	fmt.Println("customerId", c.id)
	fmt.Println("customername", c.name)
	for _, info := range c.accDetails {
		fmt.Println("account holder id is", info.Accholderid)
		fmt.Println("account number is", info.AccNo)
		fmt.Println("account balance id is", info.Balance)
		fmt.Println("account type id is", info.Acctype)
		fmt.Println("account ifsc id is", info.Accifsc)
	}
}
func (c *customer) CheckAccType(accountNo uint64, Typeacc string) *Err.MyError {
	for _, e := range c.accDetails {
		if e.AccNo == accountNo {
			if e.Acctype != Typeacc {
				return Err.New("Please Enter Valid accountType")
			}
		}
	}
	return nil
}
func (c *customer) CheckIfsc(accountNo uint64, IfscByReceiver uint64) *Err.MyError {
	for _, e := range c.accDetails {
		if e.AccNo == accountNo {
			if e.Accifsc != IfscByReceiver {
				return Err.New("invalid ifsc code")
			}
		}
	}
	return nil
}
func (c *customer) GetAccFromNo(accountNo uint64) (*acc.Account, *Err.MyError) {
	for _, e := range c.accDetails {
		if e.AccNo == accountNo {
			return e, nil
		}
	}
	return nil, Err.New("Invalid account")
}
func (c *customer) AddNewAccount(accountNo uint64, balance uint64, accounttype string, ifsc uint64) {
	c.accDetails = append(c.accDetails, acc.New(accountNo, c.id, balance, accounttype, ifsc))
}

func (c *customer) MoneyTransfer(toCustomer *customer, senderaccountNo uint64, receiveraccountno uint64, ifscByReceiver uint64, amount uint64, senderacctype string) *Err.MyError {
	senderAccount, err := c.GetAccFromNo(senderaccountNo)
	if err != nil {
		return err
	}
	err = c.CheckIfsc(senderaccountNo, ifscByReceiver)
	if err != nil {
		return err
	}
	err = c.CheckAccType(senderaccountNo, senderacctype)
	if err != nil {
		return err
	}
	if senderAccount.GetBalance() >= amount && amount > 0 {
		receiverAccount, err := toCustomer.GetAccFromNo(receiveraccountno)
		if err != nil {
			return err
		}
		receiverAccount.SetBalance(receiverAccount.GetBalance() + amount)
		senderAccount.SetBalance(senderAccount.GetBalance() - amount)
	} else {
		return Err.New("Account balance of user " + c.name + " is not sufficient")
	}
	return nil
}
