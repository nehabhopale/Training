"use strict";
exports.__esModule = true;
exports.Customer = exports.Account = void 0;
var Account = /** @class */ (function () {
    function Account(no, ifsc) {
        this.no = no;
        this.ifsc = ifsc;
    }
    return Account;
}());
exports.Account = Account;
var Bank = /** @class */ (function () {
    function Bank(name, location) {
        this.name = name;
        this.location = location;
    }
    return Bank;
}());
exports["default"] = Bank;
var Customer = /** @class */ (function () {
    function Customer(fname, lname) {
        this.fname = fname;
        this.lname = lname;
    }
    return Customer;
}());
exports.Customer = Customer;
