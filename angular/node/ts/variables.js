//Variables and methods are public by default
class ex {
    constructor() {
        this.a = 4; //by default public 
        this.b = 45;
        this.l = 23;
    }
    get variable() {
        return this.l;
    }
    set variable(value) {
        this.l = value;
    }
}
var v = new ex();
console.log(v.a);
console.log(v.b);
console.log(v.variable); //return l
v.variable = 3; //setter
console.log(v.variable);
//console.log(v.l) //error
