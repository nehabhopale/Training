class add {
    constructor(num1, num2) {
        this.num1 = num1;
        this.num2 = num2;
    }
    operation() {
        return this.num1 + this.num2;
    }
}
class sub {
    constructor(num1, num2) {
        this.num1 = num1;
        this.num2 = num2;
    }
    operation() {
        return this.num1 - this.num2;
    }
}
class mul {
    constructor(num1, num2) {
        this.num1 = num1;
        this.num2 = num2;
    }
    operation() {
        return this.num1 * this.num2;
    }
}
class div {
    constructor(num1, num2) {
        this.num1 = num1;
        this.num2 = num2;
    }
    operation() {
        return this.num1 / this.num2;
    }
}
function mathoperation(m) {
    console.log(m.operation());
}
var a = new add(2, 3);
mathoperation(a);
var a = new sub(9, 3);
mathoperation(a);
var a = new mul(2, 3);
mathoperation(a);
var a = new div(9, 3);
mathoperation(a);
