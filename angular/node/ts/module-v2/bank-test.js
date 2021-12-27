"use strict";
// import * as b from "./bank-module";
exports.__esModule = true;
// let acc=new b.Account("123","i123")
// let bank=new b.Bank("sbi","mumbai")
// let cust= new b.Customer("neha","b")
// console.log({acc,bank,cust})
/////for default class bank
var bank_module_1 = require("./bank-module");
var acc = new bank_module_1.Account("123", "i123");
var bank = new bank_module_1["default"]("sbi", "mum"); //defualt   bank 
var cust = new bank_module_1.Customer("neha", "b");
console.log({ acc: acc, bank: bank, cust: cust });
