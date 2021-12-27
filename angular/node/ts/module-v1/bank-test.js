"use strict";
exports.__esModule = true;
var bank_module_1 = require("./bank-module");
var acc = new bank_module_1.Account("123", "i123");
var bank = new bank_module_1.Bank("sbi", "mumbai");
var cust = new bank_module_1.Customer("neha", "b");
console.log({ acc: acc, bank: bank, cust: cust });
