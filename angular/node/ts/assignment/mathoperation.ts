
function add(a:number,b:number):number{
    return a+b;
}
function sub(a:number,b:number){
    return a-b;
}
function mul(a:number,b:number){
    return a*b;
}
function div(a:number,b:number):number|void{
    if (b==0){
        return console.log("b can't be zero")
       
    }
    return a/b;
}
function mathoperation(a:number,b:number,f: { (num1: number,num2:number):number|void}) {
    return f(a,b)
}
console.log(mathoperation(2,3,add))
console.log(mathoperation(4,2,sub))
console.log(mathoperation(2,3,mul))
console.log(mathoperation(9,0,div))