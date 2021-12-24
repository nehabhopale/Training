class c {
    constructor() {
        this._length = 0;
    }
    get length() {
        return this._length;
    }
    set length(value) {
        this._length = value;
    }
}
var t = new c();
t.length = 30; //this calls setter
console.log(`length is ${t.length}`); //this calls getter
