package account

import "testing"

func  TestDeposit(t *testing.T) {
	account := NewAcc(uint64(1234),"neha","n@123",uint16(200))
	
	account.Deposit(200)
	var expected uint16 = 400
	actual := account.balance
	if actual != expected {
		t.Errorf("Error found for Depositing amount")
	}
}
func  TestWithdraw(t *testing.T) {
	account := NewAcc(1234,"neha","n@123",200)
	
	account.Withdraw(100)
	var expected uint16 = 100
	actual := account.balance
	if actual != expected {
		t.Errorf("Error found for Withdrawing amount")
	}
}
func  TestAddSubscriber(t *testing.T) {
	account := NewAcc(1234,"neha","n@123",200)
	emailSub:=NewEmail("email")
	account.AddSubscriber(emailSub)
	var expected int=1
	actual := len(account.subscription)
	if actual != expected {
		t.Errorf("Error found for adding subsriber ")
	}
}

func  TestRemoveSubscriber(t *testing.T) {
	account := NewAcc(1234,"neha","n@123",200)
	emailSub:=NewEmail("email")
	account.RemoveSubscriber(emailSub)
	var expected int = 0
	actual := len(account.subscription)
	if actual != expected {
		t.Errorf("Error found for removing subscriber")
	}
}


