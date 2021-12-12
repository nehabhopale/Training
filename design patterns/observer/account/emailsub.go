package account

import "fmt"

type EmailSub struct{
	subName string
	
}

func NewEmail(subName string)EmailSub{
	return EmailSub{
		subName:subName,

	}
}

func (e EmailSub) BalanceModified(a Account){
	fmt.Println("Balance for customer",a.accUserName,"is updated for subscription",e.subName,"with balance ",a.balance)
}

