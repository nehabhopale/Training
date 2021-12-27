"use strict";
exports.__esModule = true;
var b = require("./bank-module");
var acc = new b.Account("123", "i123");
var bank = new b["default"]("sbi", "mumbai");
var cust = new b.Customer("neha", "b");
console.log({ acc: acc, bank: bank, cust: cust });
/////for default class bank
// import mybank,{Account,Customer} from "./bank-module";
// let acc =new Account("123","i123")
// let bank=new mybank("sbi","mum") //defualt   bank 
// let cust=new Customer("neha","b")
// console.log({acc,bank,cust})
///
