
// import * as b from "./bank-module";

// let acc=new b.Account("123","i123")
// let bank=new b.Bank("sbi","mumbai")
// let cust= new b.Customer("neha","b")
// console.log({acc,bank,cust})


/////for default class bank
import mybank,{Account,Customer} from "./bank-module";
let acc =new Account("123","i123")
let bank=new mybank("sbi","mum") //defualt   bank 
let cust=new Customer("neha","b")
console.log({acc,bank,cust})
 ///