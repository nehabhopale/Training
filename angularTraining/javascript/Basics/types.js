//There are 7 primitive data types: string, number, bigint, boolean, undefined, symbol, and null.

//1.string
const string1 = "A string primitive";
console.log(string1)
console.log(string1.indexOf("s"))
const string2=""
console.log(string2)
console.log(typeof string2)
const string4 = new String("A String object");//not primit it is object
console.log(string4)
//2. number
num1=Number(123)
console.log(num1)
console.log(typeof num1)

num2=123
console.log(typeof num2)
num3=12.3
console.log(typeof num3)

//3.bigint BigInt is a primitive wrapper object used to represent and 
//manipulate primitive bigint values — which are too large to be represented by the number primitive.
num1=9007199254740991
console.log(typeof num1) //number

num1=9007199254740991n
console.log(typeof num1) //bigint
num2=BigInt('1')

//4.boolean
//Any object of which the value is not undefined or null, 
//including a Boolean object whose value is false, evaluates to true
var x = new Boolean(false);
if (x) {
 console.log("inside x")///it is executed
}

var x=false
if(x){
    console.log("bye")//not executed
}

var y=!!true
console.log(y)
//If you specify any object, including a Boolean object whose value is false, 
//as the initial value of a Boolean object, the new Boolean object has a value of true.
var myFalse = new Boolean(false);   // initial value of false
console.log(myFalse)
var g = Boolean(myFalse);       // initial value of true
console.log(g)

//symbol

let c=Symbol('foo')
console.log(c)
//symbole creates new symbol at every time
console.log(Symbol('foo')==Symbol('foo'))


var x=null
console.log(typeof x)

console.log(typeof notdeclared)
