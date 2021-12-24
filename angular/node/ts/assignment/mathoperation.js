function add(a, b) {
    return a + b;
}
function sub(a, b) {
    return a - b;
}
function mul(a, b) {
    return a * b;
}
function div(a, b) {
    if (b == 0) {
        return console.log("b can't be zero");
    }
    return a / b;
}
function mathoperation(a, b, f) {
    return f(a, b);
}
console.log(mathoperation(2, 3, add));
console.log(mathoperation(4, 2, sub));
console.log(mathoperation(2, 3, mul));
console.log(mathoperation(9, 0, div));
