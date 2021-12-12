package account

type Account struct{
	accNo uint64
	accUserName string
	email string
	balance uint16
	subscription []Subscriber
}

func NewAcc(accNo uint64,accUserName string,email string,balance uint16) Account{
	return Account{
		accNo:accNo,
		accUserName:accUserName,
		email:email,
		balance:balance,
	}
}

func (a *Account)AddSubscriber(s Subscriber){
	a.subscription=append(a.subscription,s)
}

func(a *Account) RemoveSubscriber(s Subscriber){
	a.subscription=DelFromSubscription(a.subscription,s)
}

func DelFromSubscription(subList []Subscriber,subToRemove Subscriber)[]Subscriber{
	subLength:=len(subList)
	for pos ,subscription:=range subList{
		if subscription==subToRemove{
			 lastIdx:=subLength-1
			 subList[pos]=subList[lastIdx]
			 subList=subList[:pos]
			 return subList
		}
	}
	return nil
}
func (a *Account)Deposit(amt uint16){
	a.balance=a.balance+amt
}
func (a *Account)Withdraw(amt uint16){
	a.balance=a.balance-amt
}

func (a *Account)Details()[]Subscriber{
	return a.subscription
}
