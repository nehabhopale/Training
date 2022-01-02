interface mathOperation{
    num1:number;
    num2:number;
    operation():number;
}

class add{
    public num1:number
    public num2:number
    constructor(num1,num2){
        this.num1=num1;
        this.num2=num2;
    }
    operation(){
        return this.num1+this.num2;
    }
}

class sub{
    public num1:number
    public num2:number
    constructor(num1,num2){
        this.num1=num1;
        this.num2=num2;
    }
    operation(){
        return this.num1-this.num2;
    }
}
class mul{
    public num1:number
    public num2:number
    constructor(num1,num2){
        this.num1=num1;
        this.num2=num2;
    }
    operation(){
        return this.num1*this.num2;
    }
}
class div{
    public num1:number
    public num2:number
    constructor(num1,num2){
        this.num1=num1;
        this.num2=num2;
    }
    operation(){
        return this.num1/this.num2;
    }
}

function mathoperation(m:mathOperation){
    console.log(m.operation())
}


var a=new add(2,3)
mathoperation(a)
var a=new sub(9,3)
mathoperation(a)
var a=new mul(2,3)
mathoperation(a)
var a=new div(9,3)
mathoperation(a)

