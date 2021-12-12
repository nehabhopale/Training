package main

import ("obs/account"
"fmt")
//chamge in one object is notified to other
func main(){
	neha:=account.NewAcc(7678068499,"neha","nehabhopale2@gmail.com",100)

	emailSub:=account.NewEmail("email")

	neha.AddSubscriber(emailSub)
	neha.Deposit(100)
	emailSub.BalanceModified(neha)  //change in account obj(neha) is notied to subscriber obejct 
	fmt.Println(neha.Details())

	neha.RemoveSubscriber(emailSub)
	fmt.Println(neha.Details())




}