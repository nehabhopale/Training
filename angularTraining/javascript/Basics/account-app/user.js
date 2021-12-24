var account=require('./Account')
var ledger=require('./ledger')

class User{
    constructor(account,name){
        this.account=account
        this.name=name

    }
}
var acc=new account(123,1000)
var user=new User(acc,"neha")
var led=new ledger(acc)
led.listenCall()
acc.deposit(400)
acc.withdraw(10)