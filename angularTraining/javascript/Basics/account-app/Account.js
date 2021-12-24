var EventEmitter =require('events')

class Account extends EventEmitter{
    constructor(accNo,balance){
        super()
        this.accNo=accNo
        this.balance=balance
    }
    deposit(balance){
        this.balance=this.balance+balance
        this.emit('amount-deposited',this.balance,this.accNo,"deposit")
        console.log(this.balance)
    }
    withdraw(balance){
        this.balance=this.balance-balance
        this.emit('amount-withdrawan',this.balance,this.accNo,"withdraw")
        console.log(this.balance)
    }
}
module.exports=Account